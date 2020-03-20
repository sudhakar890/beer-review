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
	//Mongodb Client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://user1:user123@cluster0-9tn5d.azure.mongodb.net/test?retryWrites=true&w=majority"))
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
	beersFile, err := os.Open("sample_beers.json")
	if err != nil {
		log.Panic(err)
	}
	defer beersFile.Close()
	byteValue, _ := ioutil.ReadAll(beersFile)
	var defaultBeers []adding.Beer
	json.Unmarshal(byteValue, &defaultBeers)
	result := adding.AddBeer(connection.Beers, defaultBeers...)
	log.Println("Added sample beers data to database")

	//Adding sample data of reviews
	var beerIds []primitive.ObjectID
	for _, v := range result {
		var id = v.InsertedID.(primitive.ObjectID)
		beerIds = append(beerIds, id)
	}
	reviewsFile, err := os.Open("sample_reviews.json")
	if err != nil {
		log.Panic(err)
	}
	defer reviewsFile.Close()
	bytes, _ := ioutil.ReadAll(reviewsFile)
	var defaultReviews []adding.Review
	json.Unmarshal(bytes, &defaultReviews)
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
