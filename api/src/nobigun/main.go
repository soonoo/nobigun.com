package main

import (
	"net/http"
  "log"
  "encoding/json"
  "math/rand"

  "nobigun/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
  "github.com/go-playground/validator"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type (
  Petition struct {
    SenderName string `json:"senderName" form:"senderName" query:"senderName" validate:"required"`
    SenderAddress string `json:"senderAddress" form:"senderAddress" query:"senderAddress"`
    To string `json:"to" form:"to" query:"to"`
    Content string `json:"content" form:"content" query:"content"`
  }

  CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"http://localhost:3000", "https://nobigun.com"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))
  e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
    StackSize:  1 << 10, // 1 KB
  }))
  e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
    Format: "${time_rfc3339}, ip=${remote_ip}, latency=${latency_human}, method=${method}, uri=${uri}, status=${status}\n",
  }))

  // get recipients
  e.GET("/recipients", func(c echo.Context) error {
    recipients, err := services.GetRecipients(bson.M{})
    if err != nil {
      return c.String(http.StatusInternalServerError, "")
    }

    ret, err := json.Marshal(recipients)
    if err != nil {
      log.Println(err)
      return c.String(http.StatusInternalServerError, "")
    }

    return c.JSONBlob(http.StatusOK, ret)
  })

  // send email
  e.POST("/petitions", func(c echo.Context) error {
    p := new(Petition)
    if err := c.Bind(p); err != nil {
      return c.String(http.StatusBadRequest, err.Error())
    }
    if err := c.Validate(p); err != nil {
      return c.String(http.StatusBadRequest, err.Error())
    }

    var recipient services.Recipient
    if p.To != "" {
      objectId, err := primitive.ObjectIDFromHex(p.To)
      if err != nil {
        return err
      }
      recipients, err := services.GetRecipients(bson.M{
        "_id": objectId,
      })
      if err != nil {
        log.Println(err)
        return err
      }
      recipient = recipients[0]
    } else {
      recipients, err := services.GetRecipients(bson.M{
        "email": bson.M{ "$ne": "" },
      })
      if err != nil {
        log.Println(err)
        return err
      }
      recipient = recipients[rand.Intn(len(recipients))]
    }

    // ii, err := primitive.ObjectIDFromHex("5f38e9d8c1504a7b348e9845")
    // if err != nil {
    //   log.Println(err)
    //   return err
    // }
    // log.Println(recipient)
    go services.SendEmail(services.EmailInput{
      SenderName: p.SenderName,
      SenderEmail: p.SenderAddress,
      Content: p.Content,
      Receiver: recipient,
      // Receiver: services.Recipient{
      //   ID: ii,
      //   Email: "qpseh2m7@gmail.com",
      //   Group: "국방위원회",
      //   Name: "홍준표",
      // },
    })
    return c.String(http.StatusOK, "")
  })

  e.Logger.Debug(e.Start(":1323"))
}

