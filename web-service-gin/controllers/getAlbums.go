package controllers

import (
	album "example/web-service-gin/structs"
	utils "example/web-service-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Albums = []album.Album{
	{ID: "321", Title: "The Dark side of the moon ", Artist: "Pink Floyd", Price: 39.99},
	{ID: "567", Title: "highway to hell", Artist: "AC/DC", Price: 56.99},
	{ID: "123", Title: "Hot Fuss", Artist: "The Killers", Price: 17.99},
}

func GetAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album := utils.BinarySearch(id, Albums)
	if album.ID == id {
		c.IndentedJSON(http.StatusOK, album)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
