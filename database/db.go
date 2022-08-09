package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = ConnectDatabase()
	return Db
}

func ConnectDatabase() *gorm.DB {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp" + "(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?" + "parseTime=true&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}
	return database

}
