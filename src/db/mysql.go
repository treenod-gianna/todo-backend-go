package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var orm *gorm.DB

func OpenDatabase() *gorm.DB {
	dsn := ""
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func SetupDatabase() {
	db := OpenDatabase()
	orm = db
}

func GetORM() *gorm.DB {
	return orm
}
