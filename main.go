package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type game struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Developer string  `json:"developer"`
	Price     float64 `json:"price"`
}

// albums slice to seed record album data.
var games = []game{
	{ID: "1", Title: "Star Citizen", Developer: "CIG", Price: 45.99},
	{ID: "2", Title: "The Cycle: Frontier", Developer: "Yager", Price: 0.00},
	{ID: "3", Title: "Hell Let Loose", Developer: "Team17", Price: 29.99},
	{ID: "4", Title: "Deep Rock Galactic", Developer: "Coffe Stain Studios", Price: 19.99},
}

func main() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.GET("/games/:id", getGameById)
	router.POST("/games", postGames)
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

// postAlbums adds an album from JSON received in the request body.
func postGames(c *gin.Context) {
	var newGame game

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	// Add the new album to the slice.
	games = append(games, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getGameById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range games {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "game not found"})
}
