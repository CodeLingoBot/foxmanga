package main

import (
	_ "github.com/foxmanga/bootstrap"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	sv := gin.Default()
	sv.GET("home", func(c *gin.Context) {
		// JSON serializer is available on gin context
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})

	sv.Run(":8080")
}
