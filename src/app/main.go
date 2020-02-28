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

//Person contains id, firstname, and lastname of the person
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

//mongodb Client
var client *mongo.Client

//CreatePersonEndpoint is a post method that insert the given person into the database
func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	//Create the instance for return
	var person Person

	//Decode the request into a person object
	json.NewDecoder(request.Body).Decode(&person)

	//Connect to the database
	collection := client.Database("thepolyglotdeveloper").Collection("people")

	//Hold the resources and release once finished
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Insert the data into the database
	result, _ := collection.InsertOne(ctx, person)

	//Return the entire list once finished inserting
	json.NewEncoder(response).Encode(result)
}

//GetPeopleEndpoint gets all people in the database and returns as an array
func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	//Create the instance for return 
	var people []Person

	//Connects to the database
	collection := client.Database("thepolyglotdeveloper").Collection("people")

	//Hold the resources and release once finished
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Get all suitable solutions
	curser, err := collection.Find(ctx, bson.M{})
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":`+err.Error()+`"}`))
		return
	}
	defer curser.Close(ctx)

	//Parse each of the elements into a person object and then append back to people array
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
	
	//reencode the response and return the repsonse
	json.NewEncoder(response).Encode(people)
}

//GetPersonEndpoint returns one person from the list of people
func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	//Get the id from the url
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	//Set up the person instance
	var person Person

	//Connect to the database
	collection := client.Database("thepolyglotdeveloper").Collection("people")

	//Hold the resource and release once finished
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Find only one item according to the id in the collection, and decode the json
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":`+err.Error()+`"}`))
		return
	}
	//reencode the response and return the repsonse
	json.NewEncoder(response).Encode(person)
}

func main() {
	fmt.Println("Starting the Application...")

	//defines the boundary of the API
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Specify the mongodb connection
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	
	//Register the urls and specifies the methods of the url
	router := mux.NewRouter()
	router.HandleFunc("/person",CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people",GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}",GetPersonEndpoint).Methods("GET")

	//Register the localhost url
	http.ListenAndServe(":12345", router)
}
