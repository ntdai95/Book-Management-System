package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect(MYSQL_DB_BOOK_MANAGEMENT_SYSTEM string) {
	data, err := gorm.Open("mysql", MYSQL_DB_BOOK_MANAGEMENT_SYSTEM)
	if err != nil {
		panic(err)
	}

	db = data
}

func GetDB() *gorm.DB {
	return db
}
