package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn_orm/config"
	"learn_orm/models"
)

var (
	DB *gorm.DB
)

func InitDB() {
	config.InitConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Cfg.DB_Username,
		config.Cfg.DB_Password,
		config.Cfg.DB_Host,
		config.Cfg.DB_Port,
		config.Cfg.DB_Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		AllowGlobalUpdate: true,
	})
	if err != nil {
		panic(err)
	}

	DB = db

	DB.AutoMigrate(&models.User{}, &models.Book{})
}
