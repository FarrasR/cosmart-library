package main

import (
	"cosmart-library/database"
	migrationVersion "cosmart-library/database/migration/version"
	"cosmart-library/utils"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
)

func main() {
	utils.LoadEnv()

	ins := database.InitDB()

	m := gormigrate.New(ins.GetConn(), gormigrate.DefaultOptions, []*gormigrate.Migration{
		&migrationVersion.V20230128073015,
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration run successfully")
}
