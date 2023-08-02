package controllers

import (
	album "example/web-service-gin/structs"
	utils "example/web-service-gin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum album.Album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new album to the slice.
	Albums = append(Albums, newAlbum)
	Albums = utils.MergeSort(Albums)
	fmt.Println("new album added", newAlbum, Albums)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
