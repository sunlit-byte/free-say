package handlers

import "github.com/gin-gonic/gin"

func RegisterApiV1(r *gin.RouterGroup) {
	r.GET("/ws", SendMessageHandler)
	r.GET("/online", OnelineUserHandler)
}
