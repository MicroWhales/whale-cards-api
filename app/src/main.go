package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db, err = sql.Open("postgres", loadDBConfig())

func main() {
	router := gin.Default()
	router.GET("/", getAlbums)

	router.GET("/db/healthcheck", dbHealthCheck)
	port := os.Getenv("PORT")

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

func dbHealthCheck(c *gin.Context) {
	var response string
	healthCheckSQL := "SELECT * FROM health_check"
	row := db.QueryRow(healthCheckSQL)
	switch err := row.Scan(&response); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(response)
	default:
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, response)
}
