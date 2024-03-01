package main

import (
	"free-say/handlers"
	"free-say/room"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	go room.DefaultRoom.Run()

	v1 := router.Group("/v1")
	handlers.RegisterApiV1(v1)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	s.ListenAndServe()
}
