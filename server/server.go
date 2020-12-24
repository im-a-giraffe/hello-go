package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 * This is a basic example of starting the gin-gonic server. The server is
 * featured in a separate go package. contained in the root module.
 */
func Run() {
	fmt.Println("Starting the ginGonic server!")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
