package router

import (
	"cosmart-library/entity/response"
	"cosmart-library/handler"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer(handlers ...handler.Handler) http.Handler {
	router := gin.New()
	router.RedirectFixedPath = true

	router.GET("/", func(c *gin.Context) {
		healthCheck(c)
	})

	for _, handler := range handlers {
		handler.Register(router)
	}

	err := router.Run(os.Getenv("APP_PORT"))
	if err != nil {
		panic("Error To Start")
	}
	return router
}

func healthCheck(c *gin.Context) {
	response.OKHelloWorld(c)
}
