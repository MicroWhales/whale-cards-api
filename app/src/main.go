package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	fmt.Println("Hello Wolrd")

	router :=  gin.Default()
	router.GET("/", getAlbums)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	router.Run("localhost:" + port)
}

func getAlbums(c *gin.Context) {
	type album struct {
		ID     string  `json:"id"`
		Title  string  `json:"title"`
		Artist string  `json:"artist"`
		Price  float64 `json:"price"`
	}

	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	
    c.IndentedJSON(http.StatusOK, albums)
}