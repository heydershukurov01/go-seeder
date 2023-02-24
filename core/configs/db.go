package configs

import (
	"go.architecture/core/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Database() {
	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE")), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&models.Users{})
	DBM = db
}
