package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Album represents a simple data structure
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Database struct {
	Name string `json:"name"`
}

type Tenant struct {
	Name string `json:"name"`
}

// Sample albums slice (simulates a database)
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Kind of Blue", Artist: "Miles Davis", Price: 49.99},
}

// Get all albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func heartbeat(c *gin.Context) {
	//return nanoseconds since 1970
	c.String(http.StatusOK, "%d", time.Now().UnixNano())
	//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func version(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"version": "1.0.0"})
}

// Get album by ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func createTeanant(c *gin.Context) {
	var json Tenant
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO create tenant
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "tenant created"})
}

func getTenant(c *gin.Context) {
	name := c.Param("name")
	// TODO get tenant
	c.IndentedJSON(http.StatusOK, gin.H{"name": name})
}

func createDatabase(c *gin.Context) {
	tenant := c.Param("tenant")
	if tenant == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "tenant not found"})
		return
	}
	var json Database
	//TODO check if tenant exists
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO create database
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "database created"})
}

func getDatabase(c *gin.Context) {
	tenant := c.Param("tenant")
	if tenant == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "tenant not found"})
		return
	}
	name := c.Param("name")
	// TODO get database
	c.IndentedJSON(http.StatusOK, gin.H{"name": name})
}

const apiPrefix = "/api/v1"

// Main function
func main() {
	router := gin.Default()

	// Endpoint: GET /albums
	router.GET(apiPrefix, heartbeat)
	router.GET(apiPrefix+"/heartbeat", heartbeat)

	router.GET(apiPrefix+"/version", version)

	router.POST(apiPrefix+"/databases", createDatabase)

	router.GET(apiPrefix+"/databases/:name", getDatabase)

	router.POST(apiPrefix+"/tenants", createTeanant)

	router.GET(apiPrefix+"/tenants/:name", getTenant)

	router.Run("localhost:8080") // Start server on port 8080
}
