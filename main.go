package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nnatchy/Mee-Rai-Do/todo"
)

func main() {
	r := gin.Default();

	// Methods
	r.GET("/movies", todo.GetMovies);
	r.GET("/movies/:id", todo.GetMovie);
	r.POST("/movies", todo.InsertMovie);
	r.PUT("/movies/:id", todo.EditMovie);
	r.DELETE("/movies/:id", todo.DeleteMovie);

	r.Run(":8080");
}