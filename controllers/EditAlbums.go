package controllers

import (
	album "example/web-service-gin/structs"
	utils "example/web-service-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditAlbum(c *gin.Context) {
	var albumToEdit album.Album

	if err := c.ShouldBindJSON(&albumToEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albumEdited := false
	for i := 0; i < len(Albums); i++ {
		if utils.ConvStr(Albums[i].ID) == utils.ConvStr(albumToEdit.ID) {
			Albums[i] = albumToEdit
			albumEdited = true
			break
		}
	}

	if albumEdited {
		Albums = utils.MergeSort(Albums)
		c.IndentedJSON(http.StatusOK, "album edited")
	} else {
		c.IndentedJSON(http.StatusInternalServerError, "album has not been edited, it may not exist")
	}
}
