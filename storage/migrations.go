package storage

import (
	"log"

	"github.com/MateuszJonak/movies-api/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func Migrations() *gormigrate.Gormigrate {
	Open()

	db := GetDB()
	db.LogMode(true)
	options := &gormigrate.Options{
		TableName:    "migrations",
		IDColumnName: "id",
	}

	migrations := gormigrate.New(db, options, initialTables())

	return migrations
}

func initialTables() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "1542129194",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable(&models.User{}).Error
			},
		},
	}
}

// Migrate is used for migrating our database
func Migrate() {
	if err := Migrations().Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration did run successfully")
}

// Rollback is used for rollback our last migrations
func Rollback() {
	if err := Migrations().RollbackLast(); err != nil {
		log.Fatalf("Could not rollback: %v", err)
	}

	log.Printf("Rollback did run successfully")
}
