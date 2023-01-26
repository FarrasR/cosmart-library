package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	helloWorld     = "Hello World"
	successMessage = "Success"
)

type ResponseBody struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func buildSuccess(message string, data any) ResponseBody {
	return ResponseBody{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, buildSuccess(successMessage, data))
}

func OKWithHTTPCode(c *gin.Context, httpCode int, message string, data any) {
	c.JSON(httpCode, buildSuccess(message, data))
}

func OKHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, buildSuccess(helloWorld, nil))
}
