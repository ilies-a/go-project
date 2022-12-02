package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.New()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	server.Run(":8080")
}
