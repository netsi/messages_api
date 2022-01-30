package request

// CreateMessage defines the fields available for the POST api/v1/message request.
type CreateMessage struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Text  string `json:"text" binding:"required"`
}
