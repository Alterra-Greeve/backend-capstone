package database

import (
	"backendgreeve/config"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c config.GreeveConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DB_USER,
		c.DB_PASSWORD,
		c.DB_HOST,
		c.DB_PORT,
		c.DB_NAME,
	)

	// fmt.Println(dsn)
	//DEV MODE
	// dsn := "root:@tcp(127.0.0.1:3306)/mentalhealth?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Error("terjadi kesalahan pada database, error:", err.Error())
		return nil, err
	}

	return db, nil
}
