//Package rest provides rest api handlers
package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sudhakar890/beer-review/pkg/adding"
	"github.com/sudhakar890/beer-review/pkg/listing"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Connection of MongoDb
type Connection struct {
	Beers   *mongo.Collection
	Reviews *mongo.Collection
}

// Handler Endpoints
func Handler(client *mongo.Client) {
	//Mongo Collection
	connection := Connection{
		Beers:   client.Database("beer-review").Collection("beers"),
		Reviews: client.Database("beer-review").Collection("reviews"),
	}
	//Router
	router := mux.NewRouter()
	router.HandleFunc("/beers", connection.addBeer).Methods("POST")
	router.HandleFunc("/beers/{id}/reviews", connection.addBeerReview).Methods("POST")
	router.HandleFunc("/beers", connection.getBeers).Methods("GET")
	router.HandleFunc("/beers/{id}", connection.getBeer).Methods("GET")
	router.HandleFunc("/beers/{id}/reviews", connection.getBeerReviews).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func (connection Connection) addBeer(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var beers []adding.Beer
	err := json.NewDecoder(request.Body).Decode(&beers)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	result := adding.AddBeer(connection.Beers, beers...)
	json.NewEncoder(response).Encode(result)
}

func (connection Connection) addBeerReview(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var review adding.Review
	err := json.NewDecoder(request.Body).Decode(&review)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	_, err = listing.GetBeer(connection.Beers, id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	review.BeerID = id
	result, err := adding.AddBeerReview(connection.Reviews, review)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func (connection Connection) getBeers(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	beers, err := listing.GetBeers(connection.Beers)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(beers)
}

func (connection Connection) getBeer(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	beer, err := listing.GetBeer(connection.Beers, id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(beer)
}

func (connection Connection) getBeerReviews(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	reviews, err := listing.GetBeerReviews(connection.Reviews, id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(reviews)
}
