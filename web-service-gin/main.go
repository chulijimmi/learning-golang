package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

var albums = []album{}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameters sent by the client, then returns that album as a response
func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameters
	for _, a := range albums {
		albumID, _ := strconv.Atoi(id)
		if a.ID == albumID {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"code": "404", "message": "Album Not Found"})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/albums", getAlbums)
	r.POST("/albums/create", postAlbums)
	r.GET("/albums/:id", getAlbumsByID)
	r.Run("localhost:8080")
	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}
