package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB{
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("fail to load .env")
	}
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai",dbHost,dbUser,dbPass,dbName,dbPort)
	
	fmt.Println(dsn)
	db, errDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDB != nil {
		panic(errDB)
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB){
	dbcon, err := db.DB()
	if err != nil {
		panic("fail to close con database")
	}
	dbcon.Close()
}