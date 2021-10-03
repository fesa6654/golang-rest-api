package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func MySQLConnection() {
	con, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = con
}

func GetDB() *gorm.DB {
	return db
}
