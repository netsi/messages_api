package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"messages_api/api/message"
	"messages_api/internal/messages/loader"
	messages_repository "messages_api/internal/messages/repository"
	users_repository "messages_api/internal/users/repository"
	"os"
	"os/signal"
	"syscall"
)

const (
	apiServerPort           = 8080
	defaultMessagesFilePath = "build/data/messages.csv"
)

func main() {
	r := gin.Default()

	errChan := make(chan error)

	userRepository := users_repository.NewInMemoryUserRepository()
	messagesRepository := messages_repository.NewInMemoryRepository()

	messageLoader := loader.NewMessageLoader(messagesRepository, defaultMessagesFilePath)
	go messageLoader.Load(errChan)

	r = message.RegisterRoutes(r, userRepository, messagesRepository)

	go func() {
		err := r.Run(fmt.Sprintf(":%d", apiServerPort))
		if err != nil {
			log.Fatalf("gin.Engine.Run() returned error: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case <-quit:
		log.Println("quiting")
	case err := <-errChan:
		log.Fatalf("quiting due to error: %s", err)
	}
}
