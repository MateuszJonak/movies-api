package storage

import (
	"github.com/MateuszJonak/movies-api/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB
var logger = logrus.New()

func Open() {
	var err error
	db, err = gorm.Open(
		"mysql",
		"root:pass@/moviesdb?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil {
		logger.Error(err)
		panic("failed to connect database")
	}
	logrus.Info("Open database connection")

	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
