package main

import (
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {

	srv1 := gin.New()
	srv2 := gin.New()

	var wg sync.WaitGroup
	wg.Add(1)
	go runSrv(srv1, ":3000", "srv1")
	go runSrv(srv2, ":5000", "srv2")
	wg.Wait()
}

func runSrv(srv *gin.Engine, port string, message string) {
	srv.GET("/", func(c *gin.Context) {
		c.JSON(200, message)
	})

	srv.Run(port)
}
