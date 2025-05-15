// helper/response.go
package helper

import "strings"

// Response is the response format
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// EmptyObj is used when data is not needed
type EmptyObj struct{}

// BuildResponse builds a success response
func BuildResponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// BuildErrorResponse builds an error response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	return Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
}