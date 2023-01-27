package migrationVersion

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var V20230128090000 = gormigrate.Migration{
	ID: "20230128090000",
	Migrate: func(tx *gorm.DB) error {
		type BorrowSchedule struct {
			ReturnTime time.Time
		}

		return tx.AutoMigrate(&BorrowSchedule{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropColumn("borrow_schedules", "return_time")
	},
}
