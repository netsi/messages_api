package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"messages_api/internal/data/request"
	"messages_api/internal/data/response"
	"messages_api/internal/messages/model"
	"messages_api/internal/messages/repository"
	"net/http"
)

const (
	defaultLimit uint64 = 25
)

type handler struct {
	repository repository.Repository
	publicURL  string
}

// NewHandler returns a new instance of handler with the given dependencies
func NewHandler(repository repository.Repository, publicURL string) *handler {
	return &handler{
		repository: repository,
		publicURL:  publicURL,
	}
}

// NewMessage stores a new message.
func (h *handler) NewMessage(c *gin.Context) {
	ctx := c.Request.Context()
	req := &request.CreateMessage{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Printf("failed to bind create message request with error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response.BadRequestResponse)

		return
	}

	messageModel := model.NewMessageFromRequest(req)
	messageModel.Init()

	err = h.repository.Store(ctx, messageModel)
	if err != nil {
		log.Printf("failed to store message with error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.InternalServerErrorResponse)

		return
	}

	c.JSON(http.StatusOK, response.NewMessageResponseFromModel(messageModel))
}

// UpdateMessage text.
func (h *handler) UpdateMessage(c *gin.Context) {
	ctx := c.Request.Context()
	messageID := c.Param(MessageIDPathParameter)

	req := &request.UpdateMessage{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Printf("failed to bind update message request with error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response.BadRequestResponse)

		return
	}

	messageModel, err := h.repository.UpdateText(ctx, messageID, req.Text)
	if err != nil {
		log.Printf("failed to update message with error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.InternalServerErrorResponse)

		return
	}

	c.JSON(http.StatusOK, response.NewMessageResponseFromModel(messageModel))
}

// GetMessage from the persistence layer.
func (h *handler) GetMessage(c *gin.Context) {
	ctx := c.Request.Context()
	messageID := c.Param(MessageIDPathParameter)

	messageModel, err := h.repository.FetchByID(ctx, messageID)
	if err != nil {
		log.Printf("failed to fetch message with error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.InternalServerErrorResponse)

		return
	}

	if messageModel == nil {
		c.AbortWithStatus(http.StatusNotFound)

		return
	}

	c.JSON(http.StatusOK, response.NewMessageResponseFromModel(messageModel))
}

// GetMessages with a certain offset.
func (h *handler) GetMessages(c *gin.Context) {
	ctx := c.Request.Context()
	req := &request.GetMessages{}

	err := c.BindQuery(req)
	if err != nil {
		log.Printf("failed to bind get messages request with error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response.BadRequestResponse)

		return
	}

	messages, err := h.repository.FetchMessages(ctx, req.Offset, defaultLimit)
	if err != nil {
		log.Printf("failed to fetch messages with error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.InternalServerErrorResponse)

		return
	}

	totalItems, err := h.repository.Count(ctx)
	if err != nil {
		log.Printf("failed to fetch messages count with error: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.InternalServerErrorResponse)

		return
	}

	var nextPage string
	if totalItems >= defaultLimit && // we have more items than the limit
		req.Offset < totalItems && // the offset is smaller than the total items
		defaultLimit+req.Offset < totalItems { //the previous offset with the limit sum is lower to total count
		nextPage = fmt.Sprintf(
			"%s%s/messages?offset=%d",
			h.publicURL,
			BaseV1Api,
			req.Offset+defaultLimit,
		)
	}

	c.JSON(http.StatusOK, response.NewMessagesResponseFromModel(messages, nextPage, totalItems))
}
