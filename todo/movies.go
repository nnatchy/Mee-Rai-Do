package todo

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// ? Get all movies

func GetMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies);
}

// ? Get a movie by id

func GetMovie(c *gin.Context) {
	id := c.Param("id");
	for _, movie := range movies {
		if (movie.ID == id) {
			c.IndentedJSON(http.StatusOK, movie);
			return;
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie not found"});
}

// TODO : Insert new upcoming movie data
func InsertMovie(c *gin.Context) {
	var newMovie Movie;
	err := c.BindJSON(&newMovie);
	if (err != nil) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Structure"});
		return;
	}
	movies = append(movies, newMovie);
	
	response := gin.H{
		"message": "Upcoming movie insert successfully.",
		"movie_id": newMovie.ID,
		"movie": newMovie,
	}

	c.IndentedJSON(http.StatusOK, response);
}

// TODO : Edit upcoming movie data detail
func EditMovie(c *gin.Context) {
	id := c.Param("id");
	
	for idx, movie := range movies {
		if (movie.ID == id) {
			var updatedMovie Movie;
			movies = append(movies[:idx], movies[idx + 1:]...)
			err := c.BindJSON(&updatedMovie);
			
			if (err != nil) {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Structure"});
				return;
			}
			movies = append(movies, updatedMovie);

			response := gin.H{
				"message": "Upcoming movie update successfully.",
				"movie_id": movie.ID,
			}

			c.IndentedJSON(http.StatusOK, response);
			return;
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie ID not found"});
}

// TODO : Delete unwanted upcoming movie data
func DeleteMovie(c *gin.Context) {
	id := c.Param("id");

	for idx, movie := range movies {
		if (movie.ID == id) {

			response := gin.H{
				"message": "Upcoming movie delete successfully.",
				"movie_id": movie.ID,
			}

			movies = append(movies[:idx], movies[idx + 1:]...);
			c.IndentedJSON(http.StatusOK, response);
			return;
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie ID not found"});
}