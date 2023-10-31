package database

import (
	"fmt"
	"mhub/app/config"
	"mhub/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUSER, cfg.DBPASS, cfg.DBHOST, cfg.DBPORT, cfg.DBNAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	config.DB = db
	return db
}

func InitMigrationMysql(db *gorm.DB) {
    db.AutoMigrate(
        &models.User{},
        &models.Food{},
        &models.Order{},
		&models.OrderItem{},
        )
}
