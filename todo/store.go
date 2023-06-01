package todo

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	ID string `json: "id"`
	Name string `json: "name"`
	Release_Date string `json: "release_date"`
	Director string `json: "director"`
	Category []string `json: "category"`
}

var (
	movies []Movie
	Client *mongo.Client
)

func Init() {
	// set client to option
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017");

	// connect to mongo
	Client, err := mongo.Connect(context.TODO(), clientOptions);

	if (err != nil) {
		log.Fatal(err);
	}

	// check connection
	connectErr := Client.Ping(context.TODO(), nil);
	if (connectErr != nil) {
		log.Fatal(connectErr);
	}
	
	log.Println("Connect to Mongo successfully !");
}