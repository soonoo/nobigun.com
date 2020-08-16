package db

import (
  "time"
  "log"
  "context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func Client() (*mongo.Client, error) {
  log.Println(client)
  if client != nil {
    return client, nil
  }

  c, err := mongo.NewClient(options.Client().ApplyURI(""))
	if err != nil {
		log.Fatal(err)
    return nil, err
	}
  client = c

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
    return nil, err
	}

  err = client.Ping(ctx, readpref.Primary())
  if err != nil {
    log.Fatal(err)
    return nil, err
  }
  return client, nil
}
