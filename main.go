package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Create data
// album represnets data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//Write a hanlder to return all items
//getAlbums reponds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

//c.JSON() is able to use, too
//c.IndentedJSON is more readable
//gin.Context is the most important part of Gin.
//It carries request details, validates and serializes JSON, and more.
//Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
// http.StatusOK is 200

//Write a handler to return a specific item
//getAlbumByID locates the album whose ID value matched the id
//parameter sent by the client, then returns the album as a reponse.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

//Write a handler to add a new item
//postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	//Call BindJSON to bind the received JSON to new album.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	//Bindings are used to serialize JSON, XML, path parameters, form data, etc. to structs and maps.

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//curl -i -H "Content-type: application/json" -X POST -d "{\"id\":\"4\", \"title\":\"the modern sound of Betty Carter\", \"artist\":\"Betty Carter\", \"price\":49.99}" http://localhost:8080/albums

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
