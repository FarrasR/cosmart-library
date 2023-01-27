package database

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once

type DatabaseInstance interface {
	GetConn() *gorm.DB
	Transaction(func() error) error
}

type databaseInstance struct {
	DB *gorm.DB
}

var dbInstance *databaseInstance

func InitDB() DatabaseInstance {
	once.Do(func() {
		db, err := gorm.Open(mysql.Open(getDSN()))

		if err != nil {
			panic(err)
		}

		dbInstance = &databaseInstance{
			DB: db,
		}
	})
	return dbInstance
}

func (i *databaseInstance) GetConn() *gorm.DB {
	return i.DB
}

func (i *databaseInstance) Transaction(f func() error) error {

	i.DB = i.DB.Begin()

	err := f()
	if err != nil {
		i.DB.Rollback()
		return err
	}

	return i.DB.Commit().Error
}

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
}
