package main

import (
	"net/http"
	db "pool/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	db.ConnectDB()

	router.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
