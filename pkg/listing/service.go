//Package listing provides all list of beers and reviews
package listing

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//GetBeers returns all beers
func GetBeers(c *mongo.Collection) ([]Beer, error) {
	var ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var beers []Beer
	cursor, err := c.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var beer Beer
		cursor.Decode(&beer)
		beers = append(beers, beer)
	}
	return beers, err
}

//GetBeer returns a beer
func GetBeer(c *mongo.Collection, id primitive.ObjectID) (Beer, error) {
	var ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var beer Beer
	err := c.FindOne(ctx, bson.M{"_id": id}).Decode(&beer)
	if beer == (Beer{}) {
		err = errors.New("There is no beer for the given id")
	}
	return beer, err
}

// GetBeerReviews returns all reviews for a beer
func GetBeerReviews(c *mongo.Collection, id primitive.ObjectID) ([]Review, error) {
	var ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var reviews []Review
	cursor, err := c.Find(ctx, bson.M{"beer_id": id})
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var review Review
		cursor.Decode(&review)
		reviews = append(reviews, review)
	}
	if reviews == nil {
		err = errors.New("There is no review for the given beer id")
	}
	return reviews, err
}
