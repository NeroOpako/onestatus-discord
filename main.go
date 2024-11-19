package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		res := setup()
		c.JSON(200, res)
	})

	r.POST("/data", func(c *gin.Context) {
		var input map[string]interface{}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		c.JSON(200, gin.H{"received": input})
	})

	r.Run(":8080") // Run on port 8080
}
