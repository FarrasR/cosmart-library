package router

import (
	"cosmart-library/entity/response"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Register(router *gin.Engine)
}

func StartServer(handlers ...Handler) {
	router := gin.New()
	router.RedirectFixedPath = true

	router.GET("/", func(c *gin.Context) {
		healthCheck(c)
	})

	for _, handler := range handlers {
		handler.Register(router)
	}

	err := router.Run(":3000")
	if err != nil {
		panic("Error To Start")
	}
}

func healthCheck(c *gin.Context) {
	response.OKHelloWorld(c)
}
