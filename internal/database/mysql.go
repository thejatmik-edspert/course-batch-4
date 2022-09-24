package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateConn() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/course2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
