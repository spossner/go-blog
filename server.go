package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	srv := gin.Default()
	srv.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := srv.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
}
