package db

import (
	"database/sql"
	"fmt"

	"gin-crud2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db       *gorm.DB
	err      error
	database *sql.DB
)

func Init() {
	dsn := "root:password@tcp(localhost:3307)/app-development?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("Faied to open mysql\n%v", err.Error()))
	}

	migrate()
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	database, err = db.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to close mysql\n%v", err.Error()))
	}
	err = database.Close()
	if err != nil {
		panic(fmt.Sprintf("Failed to close mysql\n%v", err.Error()))
	}
}

func migrate() {
	db.AutoMigrate(&model.User{})
}
