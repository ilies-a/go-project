package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	srv := gin.New()

	srv.GET("/", func(c *gin.Context) {
		c.JSON(200, "srv")
	})

	srv.Run(":8080")
}
