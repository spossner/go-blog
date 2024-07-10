package main

import (
	"blog/internal/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	port := 8080

	handler := gin.Default()
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%v", port),
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       3 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           handler,
	}

	grp := handler.Group("/")
	{
		grp.GET("ping", routes.PingRoute)
	}

	err := srv.ListenAndServe() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
}
