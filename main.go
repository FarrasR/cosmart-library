package main

import (
	"cosmart-library/router"
	"cosmart-library/utils"
)

func main() {
	utils.LoadEnv()
	// database.InitDB()

	router.StartServer()

}
