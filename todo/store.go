package todo

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json: "id,omitempty"`
	Name string `bson:"name,omitempty" json: "name,omitempty"`
	Release_Date string `bson:"release_date,omitempty" json: "release_date,omitempty"`
	Director string `bson:"director,omitempty" json: "director,omitempty"`
	Category []string `bson:"category,omitempty" json: "category,omitempty"`
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