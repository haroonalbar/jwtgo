package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haroonalbar/jwtgo/initializers"
)

// this is executed before main
// init is mainly used to setup the project
func init() {
	initializers.LoadEnvVariables()
}

func main() {
	// create a router
	r := gin.Default()

	// handle ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// start server
	r.Run()
}
