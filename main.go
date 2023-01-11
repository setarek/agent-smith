package main

import (
	"agent-smith/internal/agent/handler"
	"agent-smith/internal/queue"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// handler didn't implement correctly
	// at agent client should initialize like below
	// agent := services.Agent()
	// then pass it to handler reciever to queue

	// it's better implement config, add channel's buffer size there then pass it to handler
	// we can also initialize logger in main in pass it to handler

	r.POST("/coordinate", handler.NewCoordinate)

	go queue.Walker()

	r.Run()

}
