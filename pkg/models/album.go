package models

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Col trane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarad Vaughan", Price: 33.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	rand.Seed(time.Now().UnixMicro())
	var id = rand.Intn(1000)
	var newAlbum album

	if newAlbum.ID == "" { // if the id is empty, generate a new one
		newAlbum.ID = strconv.Itoa(id) // convert int to string
	}

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatal(err)
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id") // get the id from the url

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}

func UpdateAlbumById(c *gin.Context) {
	id := c.Param("id") // get the id from the url

	for i, a := range albums { // loop through the albums
		if a.ID == id { // if the id is found

			var updateAlbum album                            // create a new album
			if err := c.BindJSON(&updateAlbum); err != nil { // bind the json to the album
				log.Fatal(err) // if there is an error, log it
			}

			albums[i] = updateAlbum                    // update the album
			c.IndentedJSON(http.StatusOK, updateAlbum) // return the updated album
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"}) // if the id is not found, return an error
}

func DeleteAlbums(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {

			albums = append(albums[:i], albums[i+1:]...)                            // delete the album
			c.IndentedJSON(http.StatusOK, gin.H{"message": "delete album success"}) // return a message
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album nor found"})
}
