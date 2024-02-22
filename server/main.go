package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Database struct {
	Name string `json:"name"`
}

type Tenant struct {
	Name string `json:"name"`
}

func heartbeat(c *gin.Context) {
	//return nanoseconds since 1970
	c.String(http.StatusOK, "%d", time.Now().UnixNano())
	//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func version(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"version": "1.0.0"})
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
