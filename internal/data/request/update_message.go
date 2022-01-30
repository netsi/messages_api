package request

// UpdateMessage defines the fields available for the PUT api/v1/message request.
type UpdateMessage struct {
	Text string `json:"text" binding:"required"`
}
