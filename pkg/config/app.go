package config

import (
	"github.com/jinzu/gorm"
	_ "github.com/jinzu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "yash:test@123/bookdata?charset=utf8&parseTime=True&local=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
