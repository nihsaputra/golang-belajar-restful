package app

import (
	"golang-belajar-restful/halper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseMysql() *gorm.DB {
	host := "localhost"
	post := "3306"
	username := "root"
	password := "@Saputra03"
	dbName := "golang_belajar_restful"

	dsn := username + ":" + password + "@tcp" + "(" + host + ":" + post + ")" + "/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	halper.PanicIfError(err)

	return db
}
