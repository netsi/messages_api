package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"messages_api/internal/auth"
	messagesRepository "messages_api/internal/messages/repository"
	usersRepository "messages_api/internal/users/repository"
)

const (
	MessageIDPathParameter = "message_id"
	BaseV1Api              = "/api/v1"
	defaultURL             = "http://localhost:8080"
)

// RegisterRoutes adds the message API routes available to the gin.Engine
func RegisterRoutes(
	r *gin.Engine,
	userRepository usersRepository.Repository,
	messagesRepository messagesRepository.Repository,
) *gin.Engine {
	messageHandler := NewHandler(messagesRepository, defaultURL)

	publicAPI := r.Group(BaseV1Api)
	publicAPI.POST("/message", messageHandler.NewMessage)

	privateAPI := r.Group(BaseV1Api)
	privateAPI.Use(auth.BasicAuth(userRepository))
	privateAPI.PUT(fmt.Sprintf("/message/:%s", MessageIDPathParameter), messageHandler.UpdateMessage)
	privateAPI.GET(fmt.Sprintf("/message/:%s", MessageIDPathParameter), messageHandler.GetMessage)
	privateAPI.GET("/messages", messageHandler.GetMessages)

	return r
}
