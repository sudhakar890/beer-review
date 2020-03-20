package main

import (
	"context"
	"log"
	"time"

	"github.com/sudhakar890/beer-review/pkg/http/rest"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//Mongodb Client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://<user>:<password>@cluster0-9tn5d.azure.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Panic(err)
	}
	// Context for timeout of 5 secs
	var ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(ctx)

	//Start the application
	rest.Handler(client)
}
