package storage

import (
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Open() {
	var err error
	db, err = gorm.Open(
		"mysql",
		"root:pass@/moviesdb?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}
	log.Printf("Open database connection")
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
