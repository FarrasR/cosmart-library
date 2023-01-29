package migrationVersion

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var V20230129090000 = gormigrate.Migration{
	ID: "20230129090000",
	Migrate: func(tx *gorm.DB) error {
		type Book struct {
			gorm.Model
			Title   string `gorm:"type:varchar(255);not null"`
			Author  string `gorm:"type:varchar(255);not null"`
			Genre   string `gorm:"type:varchar(255);not null"`
			Edition int
		}

		return tx.AutoMigrate(&Book{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropColumn("books", "genre")
	},
}
