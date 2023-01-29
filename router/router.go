package router

import (
	"cosmart-library/entity/response"
	"cosmart-library/handler"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func BuildHandler(handlers ...handler.Handler) *gin.Engine {
	router := gin.New()

	for _, handler := range handlers {
		handler.Register(router)
	}
	return router
}

func RunServer(router *gin.Engine) {
	router.RedirectFixedPath = true
	router.HandleMethodNotAllowed = true

	router.GET("/", func(c *gin.Context) {
		healthCheck(c)
	})
	router.NoRoute(noRoute)
	router.NoMethod(noMethod)
	router.Use(handleInternalError)

	err := router.Run(os.Getenv("APP_PORT"))
	if err != nil {
		panic("Error To Start")
	}
}

func healthCheck(c *gin.Context) {
	response.OKHelloWorld(c)
}

func noRoute(c *gin.Context) {
	response.Error(c, http.StatusNotFound, "Not Found")
}

func noMethod(c *gin.Context) {
	response.Error(c, http.StatusMethodNotAllowed, "Method Not Allowed")

}

func handleInternalError(c *gin.Context) {
	if len(c.Errors) <= 0 {
		c.Next()
		return
	}

	response.Error(c, http.StatusInternalServerError, c.Errors.Last().Error())
}
