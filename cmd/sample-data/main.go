package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/sudhakar890/beer-review/pkg/adding"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Connection of MongoDb
type Connection struct {
	Beers   *mongo.Collection
	Reviews *mongo.Collection
}

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
		log.Panic(err)
	}
	// Context for timeout of 5 secs
	var ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(ctx)
	//Mongo Collection
	connection := Connection{
		Beers:   client.Database("beer-review").Collection("beers"),
		Reviews: client.Database("beer-review").Collection("reviews"),
	}

	//Adding sample data of beers
	beersFile, err := ioutil.ReadFile("/data/sample_beers.json")
	if err != nil {
		log.Panic(err)
	}
	var defaultBeers []adding.Beer
	json.Unmarshal(beersFile, &defaultBeers)
	result := adding.AddBeer(connection.Beers, defaultBeers...)
	log.Println("Added sample beers data to database")

	//Adding sample data of reviews
	var beerIds []primitive.ObjectID
	for _, v := range result {
		var id = v.InsertedID.(primitive.ObjectID)
		beerIds = append(beerIds, id)
	}
	reviewsFile, err := ioutil.ReadFile("/data/sample_reviews.json")
	if err != nil {
		log.Panic(err)
	}
	var defaultReviews []adding.Review
	json.Unmarshal(reviewsFile, &defaultReviews)
	for i := range defaultReviews {
		if i%2 == 0 {
			defaultReviews[i].BeerID = beerIds[0]
			adding.AddBeerReview(connection.Reviews, defaultReviews[i])
		} else {
			defaultReviews[i].BeerID = beerIds[1]
			adding.AddBeerReview(connection.Reviews, defaultReviews[i])
		}
	}
	log.Println("Added sample reviews data to database")
}
