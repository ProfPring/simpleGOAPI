package main

import (
	controllers "example/web-service-gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/albums", controllers.GetAllAlbums)
	router.POST("/postAlbums", controllers.PostAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.DELETE("/DeleteAlbum/:id", controllers.DeleteAlbum)
	router.PUT("/UpdateAlbum", controllers.EditAlbum)

	router.Run("localhost:8080")
}
