package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		// for i := 0; i < 1000000000; i++ {
		// }
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.Run("0.0.0.0:8080")
}
