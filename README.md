# RESTful-API-with-Go-and-Gin

ref: 
- [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)
- [Gin Web Framework package documentation](https://pkg.go.dev/github.com/gin-gonic/gin) 
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
 
---
![](https://i.imgur.com/YfZuPsq.png)


## Import package and Create the data
```go
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//struct
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// slice
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
```

## Write a hanlder to return all items

```go
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
```

- Note that you can replace `Context.IndentedJSON` with a call to `Context.JSON` to send more compact JSON. In practice, the indented form is much easier to work with when debugging and the size difference is usually small.


### Update main function
```go
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)

    router.Run("localhost:8080")
}
```

|Test on Postman|
|-|
|![](https://i.imgur.com/8I6TivF.png)|



## Write a handler to return a specific item
```go
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
```

### Update main function
```go
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}
```

|Test on Postman|
|-|
|![](https://i.imgur.com/fcPiANW.png)|

## Write a handler to add a new item
```go
func postAlbums(c *gin.Context) {
	var newAlbum album

	//Call BindJSON to bind the received JSON to new album.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	
	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
```
- Bindings are used to serialize JSON, XML, path parameters, form data, etc. to structs and maps

### Update main function
```go
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
```

|Test on Postman|
|-|
|![](https://i.imgur.com/nVMVvae.png)|
|![](https://i.imgur.com/T5qAtr2.png)|
