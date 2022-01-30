package response

// ErrorResponse returned when an error occurs
type ErrorResponse struct {
	Message string `json:"message"`
}

var BadRequestResponse = ErrorResponse{Message: "invalid api call"}
var InternalServerErrorResponse = ErrorResponse{Message: "unexpected error, please try again later"}
