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
	borrowScheduleRepository := repository.NewBorrowScheduleRepository(database)

	bookService := service.NewBookService(bookRepository)
	borrowScheduleService := service.NewBorrowScheduleService(borrowScheduleRepository)

	router.StartServer(
		handler.NewBookHandler(bookService),
		handler.NewBorrowScheduleHandler(borrowScheduleService),
	)
}
