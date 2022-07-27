package main

import (
	"example/web-service/pkg/models"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	router.GET("/albums", models.GetAlbums)
	router.POST("/albums", models.PostAlbum)
	router.GET("/albums/:id", models.GetAlbumById)
	router.PUT("/albums/:id", models.UpdateAlbumById)
	router.DELETE("/albums/:id", models.DeleteAlbums)
	router.Run("0.0.0.0" + port)
}
