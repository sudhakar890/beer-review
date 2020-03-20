//Package adding provides beer adding operations
package adding

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//AddBeer adds the given beer(s) to the database
func AddBeer(c *mongo.Collection, b ...Beer) []*mongo.InsertOneResult {
	var results []*mongo.InsertOneResult
	// Context for timeout of 5 secs
	var ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	for _, beer := range b {
		result, err := c.InsertOne(ctx, beer)
		if err != nil {
			log.Panic(err)
		}
		results = append(results, result)
	}
	return results
}

// AddBeerReview saves a new beer review in the database
func AddBeerReview(c *mongo.Collection, r Review) (*mongo.InsertOneResult, error) {
	r.Created = time.Now()
	// Context for timeout of 5 secs
	var ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	result, err := c.InsertOne(ctx, r)
	return result, err
}
