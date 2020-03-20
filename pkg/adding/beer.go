package adding

//Beer defines properties of beer to be added
type Beer struct {
	Name      string  `json:"name" bson:"name"`
	Brewery   string  `json:"brewery" bson:"brewery"`
	Abv       float32 `json:"abv" bson:"abv"`
	ShortDesc string  `json:"short_description" bson:"short_description"`
}
