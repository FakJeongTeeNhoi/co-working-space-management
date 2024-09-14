package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenericJSON map[string]interface{}

type ErrorResponse struct {
	Status  int    `json:"errorCode"`
	Message string `json:"message,omitempty"`
}

type CommonResponse struct {
	Success bool `json:"success"`
}

// success response
func (c CommonResponse) AddInterfaces(value map[string]interface{}) map[string]interface{} {
	var response = make(map[string]interface{})
	response["success"] = c.Success
	for key, value := range value {
		response[key] = value
	}
	return response
}

func Errorf(status int, message string, args ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:  status,
		Message: fmt.Sprintf(message, args...),
	}
}

func (e ErrorResponse) Error() string {
	return fmt.Sprint("%d:%s", e.Status, e.Message)
}

func (e ErrorResponse) AbortWithError(c *gin.Context) {
	c.AbortWithStatusJSON(e.Status, e)
}

func Unauthorized(message string) *ErrorResponse {
	return Errorf(http.StatusUnauthorized, message)
}

func BadRequest(message string) *ErrorResponse {
	return Errorf(http.StatusBadRequest, message)
}

func InternalServerError(message string) *ErrorResponse {
	return Errorf(http.StatusInternalServerError, message)
}

func Forbidden(message string) *ErrorResponse {
	return Errorf(http.StatusForbidden, message)
}

func NotFound(message string) *ErrorResponse {
	return Errorf(http.StatusNotFound, message)
}
