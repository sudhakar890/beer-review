package listing

import "go.mongodb.org/mongo-driver/bson/primitive"

// Beer defines the properties of a beer to be listed
type Beer struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Brewery   string             `json:"brewery" bson:"brewery"`
	Abv       float32            `json:"abv" bson:"abv"`
	ShortDesc string             `json:"short_description" bson:"short_description"`
}
