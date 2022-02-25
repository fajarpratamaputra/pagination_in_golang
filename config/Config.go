package config

import (
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (database *gorm.DB, err error) {
	var dsn strings.Builder
	dsn.WriteString(os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")" + "/" + os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local")
	database, err = gorm.Open(mysql.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return database, nil
}
