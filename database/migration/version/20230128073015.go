package migrationVersion

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var V20230128073015 = gormigrate.Migration{
	ID: "V20230128073015",
	Migrate: func(tx *gorm.DB) error {
		type Book struct {
			gorm.Model
			Title   string `gorm:"type:varchar(255);not null"`
			Author  string `gorm:"type:varchar(255);not null"`
			Edition int
		}
		type BorrowSchedule struct {
			gorm.Model
			Name       string `gorm:"type:varchar(255);not null"`
			BookId     int    `gorm:"index"`
			PickupTime time.Time
			DueTime    time.Time
		}

		return tx.AutoMigrate(&Book{}, &BorrowSchedule{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("book", "borrow_schedule")
	},
}
