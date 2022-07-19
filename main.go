package main

import (
	"example/web-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", models.GetAlbums)
	router.POST("/albums", models.PostAlbum)
	router.GET("/albums/:id", models.GetAlbumById)
	router.PUT("/albums/:id", models.UpdateAlbumById)
	router.DELETE("/albums/:id", models.DeleteAlbums)
	router.Run("localhost:8080")
}
