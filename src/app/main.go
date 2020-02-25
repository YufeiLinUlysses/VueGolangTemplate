package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Person contains person
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

//mongodb Client
var client *mongo.Client

//CreatePersonEndpoint creates
func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

//GetPeopleEndpoint gets
func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var people []Person
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	curser, err := collection.Find(ctx, bson.M{})
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":`+err.Error()+`"}`))
		return
	}
	defer curser.Close(ctx)
	for curser.Next(ctx){
		var person Person
		curser.Decode(&person)
		people = append(people, person)
	}
	if err := curser.Err(); err!=nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":`+err.Error()+`"}`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

//GetPersonEndpoint creates
func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":`+err.Error()+`"}`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func main() {
	fmt.Println("Starting the Application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	router := mux.NewRouter()
	router.HandleFunc("/person",CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people",GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}",GetPersonEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}
