package router

import (
	"cosmart-library/entity/response"
	"cosmart-library/handler"
	"os"

	"github.com/gin-gonic/gin"
)

func BuildHandler(handlers ...handler.Handler) *gin.Engine {
	router := gin.New()
	router.RedirectFixedPath = true

	for _, handler := range handlers {
		handler.Register(router)
	}
	return router
}

func RunServer(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		healthCheck(c)
	})

	err := router.Run(os.Getenv("APP_PORT"))
	if err != nil {
		panic("Error To Start")
	}
}

func healthCheck(c *gin.Context) {
	response.OKHelloWorld(c)
}
