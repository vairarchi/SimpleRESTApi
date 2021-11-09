package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Defining connection
func Connect() {
	// You can define your user name and password for my sql.
	// d, err := gorm.Open("mysql", "root:root@/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", "root:@(localhost)/book?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// Gets DB
func GetDB() *gorm.DB {
	return db
}
