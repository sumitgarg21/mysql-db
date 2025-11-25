package config

import(
	"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB * gorm.DB
)

func Connect(){
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",user,pass,dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err!=nil{
		panic(err)
	}
	DB = db
}

func GetDB() *gorm.DB{
	return DB
}