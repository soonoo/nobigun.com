package services

import (
  "context"
  "time"
  "log"
  "fmt"
  // "os"
  // "bufio"
  // "strings"
  "nobigun/db"

  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/bson"
  "github.com/aws/aws-sdk-go/service/ses"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
)

type (
  Recipient struct {
    ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
    Email string `json:"email"`
    Group string `json:"group"`
    Name string `json:"name"`
  }
  EmailInput struct {
    SenderName string
    SenderEmail string
    Content string
    Receiver Recipient
  }
)

func GetRecipients(filter interface{}) ([]Recipient, error) {
  client, err := db.Client()
  if err != nil {
    log.Println(err)
    return nil, err
  }

  db := client.Database("nobigun")
  recipientsCollection := db.Collection("recipients")

  ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
  cursor, err := recipientsCollection.Find(ctx, filter)
  if err != nil {
    log.Println(err)
    return nil, err
  }

  var recipients []Recipient
  if err = cursor.All(ctx, &recipients); err != nil {
    log.Println(err)
    return nil, err
  }

  return recipients, nil
}

func SendEmail(emailInput EmailInput) {
  client, err := db.Client()
  if err != nil {
    log.Println(err)
    return
  }

  db := client.Database("nobigun")
  statCollection := db.Collection("stat")

  sess, err := session.NewSession(&aws.Config{
    Region:      aws.String("ap-northeast-2"),
    Credentials: credentials.NewSharedCredentials("", "nobigun"),
  })
  if err != nil {
    log.Println(err)
    return
  }

  userMessage := ""
  if emailInput.Content != "" {
    userMessage = fmt.Sprintf(
`%s 님의 메세지:
%s
`,
emailInput.SenderName, emailInput.Content)
  }

  formattedContent := fmt.Sprintf(
`안녕하세요 %s 의원님.
(이 이메일은 대한민국 유권자인 %s님을 대신하여 nobigun.com에서 발송되었습니다.)

수도권을 중심으로 코로나19 확진자가 연일 폭증하고 있는 가운데 국방부는 2020년 9월 1일부터 예비군 훈련을 재개한다고 발표하였습니다.

사회적 거리두기 2단계 격상과 함께 전국민이 코로나19의 공포에 떨고 있습니다.
코로나19 사태가 진정되지 않은 현 상황에서 예비군 훈련을 강행할 경우, 300만 예비군은 물론 5,000만 국민의 건강은 그 누구도 보장할 수 없습니다.
훈련 시간 단축은 대안이 아닙니다. 코로나19 사태로부터 안전을 보장받을 수 있는 방법은 예비군 훈련을 취소하는것 뿐입니다.


%s 소속이신 %s 의원님께서 국민들의 안전을 위해 목소리를 내어주십시오.
답장 기다리겠습니다.

%s
`,
emailInput.Receiver.Name, emailInput.SenderName, emailInput.Receiver.Group, emailInput.Receiver.Name, userMessage)

  ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
  _, err = statCollection.InsertOne(ctx, bson.D{
    {Key: "receiver", Value: emailInput.Receiver.Name},
    {Key: "receiverAddress", Value: emailInput.Receiver.Email},
    {Key: "content", Value: formattedContent},
    {Key: "senderAddress", Value: emailInput.SenderEmail},
    {Key: "senderName", Value: emailInput.SenderName},
  })
  if err != nil {
    log.Println(err)
  }

  svc := ses.New(sess)
  var ccList []*string
  if emailInput.SenderEmail != "" {
    ccList = append(ccList, aws.String(emailInput.SenderEmail))
  }

  log.Println(emailInput.Receiver.Email)
  _, err = svc.SendEmail(&ses.SendEmailInput{
    Message: &ses.Message{
      Body: &ses.Body{
        Text: &ses.Content{
          Data: aws.String(formattedContent),
        },
      },
      Subject: &ses.Content{
        Data: aws.String("의원님, 코로나19 사태 속 300만 예비군 훈련 강행에 대한 입장이 궁금합니다."),
      },
    },
    Destination: &ses.Destination{
      ToAddresses: []*string{
        aws.String(emailInput.Receiver.Email),
      },
      CcAddresses: ccList,
    },
    Source: aws.String("nobi@nobigun.com"),
  })
  if err != nil {
    log.Println(err)
  }
  return
}

	// ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
  // f, err := os.Open("services/members")
  // if err != nil {
    // log.Println(err)
  // }
  // scanner := bufio.NewScanner(f)
  // log.Println(f)
  // for scanner.Scan() {
    // line := scanner.Text()
    // arr := strings.Split(line, " ")
    // email := ""
    // log.Println(line)
    // if len(arr) == 3 {
      // email = arr[2]
    // }
    // recipientsCollection.InsertOne(ctx, bson.D{
      // {Key: "email", Value: email},
      // {Key: "name", Value: arr[0]},
      // {Key: "group", Value: arr[1]},
    // })
  // }

