package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getAlbums)
	loadDBConfig()
	port := os.Getenv("PORT")
	// db := os.Getenv("DB_USERNAME")
	var address string = ""

	if port == "" {
		// if port environmental variable isn't specified
		address = "127.0.0.1:5000"
	} else {
		address = "0.0.0.0:5000"
	}

	router.Run(address)
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
