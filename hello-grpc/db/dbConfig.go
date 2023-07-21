package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDb() {

	dsn := "root:951004Lx@tcp(127.0.0.1:33360)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN: dsn,
		}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

}
