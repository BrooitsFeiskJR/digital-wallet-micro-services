package responses

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	StatusCode string    `json:"status_code"`
	Data       any       `json:"data,omitempty"`
	Message    string    `json:"message,omitempty"`
	Links      any       `json:"_links"`
	Error      string    `json:"error,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Type string `json:"type,omitempty"`
}

func SuccessResponse(data any, statusCode int, message ...string) *APIResponse {
	response := &APIResponse{
		StatusCode: strconv.Itoa(statusCode),
		Data:       data,
		Timestamp:  time.Now().UTC(),
	}

	if len(message) > 0 {
		response.Message = message[0]
	}

	return response
}

func ErrorResponse(statusCode int, errorMessage string) *APIResponse {
	return &APIResponse{
		StatusCode: strconv.Itoa(statusCode),
		Error:      http.StatusText(statusCode),
		Message:    errorMessage,
		Timestamp:  time.Now().UTC(),
	}
}

func ValidationError(statusCode string, errors map[string]string) *APIResponse {
	return &APIResponse{
		StatusCode: statusCode,
		Error:      "Validation Error",
		Data:       errors,
		Timestamp:  time.Now().UTC(),
	}
}

func (res *APIResponse) ToJSON(c *gin.Context, statusCode int) {
	c.JSON(statusCode, res)
	c.Abort()
}
