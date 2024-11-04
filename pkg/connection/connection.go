package connection

import (
	"Blog_API/pkg/config"
	"Blog_API/pkg/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// Connect to the database
func Connect() {
	dbConfig := config.LocalConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBIp, dbConfig.DBName)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Error connecting to DB")
		panic(err)
	}
	fmt.Println("Database Connected")
	db = d
}

// Creating New table in foodstore database
func Migrate() {
	db.Migrator().AutoMigrate(models.User{})
	db.Migrator().AutoMigrate(models.BlogPost{})
	db.Migrator().AutoMigrate(models.Comment{})
	db.Migrator().AutoMigrate(models.Reaction{})
}

// Calling to connect function to initalize connection
func GetDB() *gorm.DB {
	if db == nil {
		Connect()
	}
	Migrate()
	return db
}
