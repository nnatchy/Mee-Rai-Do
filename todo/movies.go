package todo

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ? Get all movies

func GetMovies(c *gin.Context) {
	var movies []Movie;
	collections := Client.Database("mee-rai-do").Collection("movies");

	cur, _ := collections.Find(context.Background(), bson.D{});

	for cur.Next(context.Background()) {
		var movie Movie;
		err := c.BindJSON(&movie);
		if (err != nil) {
			log.Fatal(err);
		}
		movies = append(movies, movie);
	}
	err := cur.Err();
	if (err != nil) {
		log.Fatal(err);
	}
	cur.Close(context.Background());
	
	c.IndentedJSON(http.StatusOK, movies);
}

// ? Get a movie by id

func GetMovie(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"));
	if (err != nil) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"});
		return;
	}
	movieCollection := Client.Database("mee-rai-do").Collection("movies");
	var movie Movie;
	var findErr = movieCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&movie);
	if (findErr != nil) {
		if (findErr == mongo.ErrNoDocuments) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie not found"});
			return;
		}
		// Server error
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"});
	}
	c.IndentedJSON(http.StatusOK, movie);
}

// TODO : Insert new upcoming movie data
func InsertMovie(c *gin.Context) {
	var newMovie Movie;
	err := c.BindJSON(&newMovie);
	if (err != nil) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Structure"});
		return;
	}
	moviesCollection := Client.Database("mee-rai-do").Collection("movies");
	insertRes, err := moviesCollection.InsertOne(context.TODO(), newMovie);
	
	if (err != nil) {
		c.IndentedJSON(http.StatusInternalServerError, err.Error());
		return;
	}

	response := gin.H{
		"message": "Upcoming movie insert successfully.",
		"movie_id": newMovie.ID,
		"movie": newMovie,
	}
	newMovie.ID = insertRes.InsertedID.(primitive.ObjectID)
	c.IndentedJSON(http.StatusOK, response);
}

// TODO : Edit upcoming movie data detail
func EditMovie(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"));
	if (err != nil) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"});
		return;
	}
	moviesCollection := Client.Database("mee-rai-do").Collection("movies");

	var jsonMovie Movie;
	bindErr := c.BindJSON(&jsonMovie);
	if (bindErr != nil) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Structure"});
		return;
	}
	
	updatedMovie := bson.D {
		{"$set", bson.D {
			{"name", jsonMovie.Name},
			{"release_date", jsonMovie.Release_Date},
			{"director", jsonMovie.Director},
			{"category", jsonMovie.Category},
		}},
	}

	res, err := moviesCollection.UpdateOne(context.Background(), bson.M{"_id": id}, updatedMovie);

	// Server Error
	if (err != nil) {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Server Error"});
		return;
	}

	if (res.MatchedCount == 0) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Movie Found"});
		return;
	}

	response := gin.H{
		"message": "Upcoming Movie Data updated successfully",
		"movie_id": id,
	}
	c.IndentedJSON(http.StatusNotFound, response);
}

// TODO : Delete unwanted upcoming movie data
func DeleteMovie(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"));
	if (err != nil) {
		c.IndentedJSON(http.StatusBadRequest, err.Error());
		return;
	}

	moviesCollection := Client.Database("mee-rai-do").Collection("movies");
	var deleteErr = moviesCollection.FindOneAndDelete(context.Background(), bson.M{"_id": id});
	if (deleteErr != nil) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie ID not found"});
		return;
	}
	response := gin.H{
		"message": "Upcoming movie delete successful",
		"movie_id": id,
	}
	c.IndentedJSON(http.StatusOK, response);
}