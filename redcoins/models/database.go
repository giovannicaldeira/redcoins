package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"fmt"
	"os"
	"time"
)

var db *gorm.DB

func init() {

	env := godotenv.Load()
	if env != nil {
		fmt.Print(env)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURI := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True", username, password, dbHost, dbPort, dbName)
	fmt.Println(dbURI)

	conn, err := gorm.Open("mysql", dbURI)
	for {
		if err != nil {
			conn, err = gorm.Open("mysql", dbURI)
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}

	db = conn
	db.Debug().AutoMigrate(&Operation{}, &User{}, &Bitcoin{})
	db.Model(&Operation{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}

func GetDB() *gorm.DB {
	return db
}
