package model

import (
	"time"

	"gorm.io/gorm"
)

type BorrowSchedule struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255);not null"`
	BookId     int    `gorm:"index"`
	PickupTime *time.Time
	DueTime    *time.Time
	ReturnTime *time.Time
}
