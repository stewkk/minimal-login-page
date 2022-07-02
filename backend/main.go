package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type entry struct {
	ID string `json:"id"`
	Value string `json:"value"`
}

var entries = []entry {
	{ID: "1", Value: "12345"},
	{ID: "2", Value: "Vyacheslav, где фронт?"},
	{ID: "351", Value: "pukpuk"},
}

func main() {
	router := gin.Default()
	router.GET("/entries", getEntries)

	router.Run("localhost:8080")
}

func getEntries(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, entries)
}

