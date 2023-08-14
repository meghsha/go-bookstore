package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func connect() {
	d, err := gorm.Open("mysql", "root:password/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

func getDatabase() *gorm.DB {
	return db
}
