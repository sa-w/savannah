package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Database utilities

func ConnectDatabase() *gorm.DB {

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db

}
