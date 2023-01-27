package main

import (
	"cosmart-library/database"
	"cosmart-library/router"
	"cosmart-library/utils"
)

func main() {
	utils.LoadEnv()
	database.InitDB()

	database.GetConn()

	router.StartServer()

}
