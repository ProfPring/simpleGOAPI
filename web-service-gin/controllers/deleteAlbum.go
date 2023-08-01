package controllers

import (
	utils "example/web-service-gin/utils"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(Albums)

	Albums = utils.DeleteAlbum(id, Albums)

	fmt.Println("albums after delete", Albums)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
}
