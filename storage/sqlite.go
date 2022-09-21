package storage

import (
	"sync"

	"github.com/Aakash-Pandit/ADmyBRAND/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var once sync.Once
var Database *gorm.DB

func GetDatabase() *gorm.DB {
	return Database
}

func NewConnection() (*gorm.DB, error) {
	once.Do(
		func() {
			var err error
			Database, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
			if err != nil {
				panic("failed to connect Database")
			}

			err = Database.AutoMigrate(&models.User{})
			if err != nil {
				panic("failed to migrate Database")
			}
		})

	return Database, nil
}
