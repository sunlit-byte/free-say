package handlers

import "github.com/gin-gonic/gin"

func RegisterApiV1(r *gin.RouterGroup) {
	r.POST("/ws", SendMessageHandler)
}
