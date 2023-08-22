package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Engine *gorm.DB

func InitDB(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Engine = db
	s, err := Engine.DB()
	s.SetMaxOpenConns(50)
	s.SetMaxIdleConns(10)
	s.SetConnMaxLifetime(time.Hour)
}
