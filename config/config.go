package config

import (
	"log"
	"mhub/models"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func InitializeDatabase() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbConfig := DatabaseConfig{
		Host: os.Getenv("DB_HOST"),
	}
	portStr := os.Getenv("DB_PORT")
	dbConfig.Port, _ = strconv.Atoi(portStr)
	dbConfig.Username = os.Getenv("DB_USERNAME")
	dbConfig.Password = os.Getenv("DB_PASSWORD")
	dbConfig.DBName = os.Getenv("DB_NAME")

	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" + dbConfig.DBName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Food{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.OrderDetail{})

	return db, nil
}
