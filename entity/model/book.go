package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null"`
	Author  string `gorm:"type:varchar(255);not null"`
	Edition int
	Genre   string `gorm:"type:varchar(255);not null"`
}
