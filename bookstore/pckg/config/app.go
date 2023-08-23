package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// whole point of this file is to return a DB so that other files can access the db
var (
	db *gorm.DB
)

func Conect() {
	d, err := gorm.Open("mysql", "user:password/table?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
