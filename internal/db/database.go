package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "aetherGo.db"
	}

	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection established")

}

func AutoMigrate(models ...interface{}) {
	for _, model := range models {
		if !DB.Migrator().HasTable(model) {
			log.Printf("Migrating table for model: %T", model)
			err := DB.AutoMigrate(model)
			if err != nil {
				log.Fatalf("Failed to migrate model %T: %v", model, err)
			} else {
				log.Printf("Migration successful for model: %T", model)
			}
		} else {
			log.Printf("Model %T already migrated, skipping.", model)
		}
	}
}
