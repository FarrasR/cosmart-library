package main

import (
	"cosmart-library/database"
	"cosmart-library/handler"
	"cosmart-library/repository"
	"cosmart-library/router"
	"cosmart-library/service"
	"cosmart-library/utils"
)

func main() {
	utils.LoadEnv()
	database := database.InitDB()
	bookRepository := repository.NewBookRepository(database)
	bookService := service.NewBookService(bookRepository)
	BookHandler := handler.NewBookHandler(bookService)

	router.StartServer(BookHandler)
}
