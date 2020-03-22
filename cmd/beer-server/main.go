package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/sudhakar890/beer-review/pkg/http/rest"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//Env
	var dbUser string
	var dbPassword string
	if os.Getenv("DB_USER") != "" {
		dbUser = os.Getenv("DB_USER")
	}
	if os.Getenv("DB_PASSWORD") != "" {
		dbPassword = os.Getenv("DB_PASSWORD")
	}
	if os.Getenv("DB_USER_FILE") != "" {
		file := os.Getenv("DB_USER_FILE")
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Panic(err)
		}
		dbUser = string(bytes)
	}
	if os.Getenv("DB_PASSWORD_FILE") != "" {
		file := os.Getenv("DB_PASSWORD_FILE")
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Panic(err)
		}
		dbPassword = string(bytes)
	}
	//Mongodb Client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + dbUser + ":" + dbPassword + "@mongodb:27017"))
	if err != nil {
		log.Fatal(err)
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
