package database

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB
var once sync.Once

func InitDB() {
	once.Do(func() {
		var err error
		dbInstance, err = gorm.Open(mysql.Open(GetDSN()))
		if err != nil {
			panic(err)
		}
	})
}

func GetConn() *gorm.DB {
	return dbInstance
}

func Transaction(f func() error) error {

	dbInstance = dbInstance.Begin()

	err := f()
	if err != nil {
		dbInstance.Rollback()
		return err
	}

	return dbInstance.Commit().Error
}

func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
}
