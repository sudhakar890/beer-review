package listing

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Review defines a beer review
type Review struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	BeerID    primitive.ObjectID `json:"beer_id" bson:"beer_id"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Score     int                `json:"score" bson:"score"`
	Text      string             `json:"text" bson:"text"`
	Created   time.Time          `json:"created" bson:"created"`
}
