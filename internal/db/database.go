package db

import (
	"AetherGo/internal/log"
	"os"

	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(appName string) {
	var err error

	if err := os.MkdirAll(appName, 0755); err != nil {
		log.Fatalf("Failed to create app directory: %v", err)
	}

	dsn := filepath.Join(appName, "aether.sqlite")

	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := DB.Exec("PRAGMA journal_mode=DELETE;").Error; err != nil {
		log.Infof("Warning: Failed to set journal mode: %v", err)
	}
}

func GetDB() *gorm.DB {
	return DB
}

func AutoMigrate(models ...interface{}) {
	for _, model := range models {
		if !DB.Migrator().HasTable(model) {
			err := DB.AutoMigrate(model)
			if err != nil {
				log.Fatalf("Failed to migrate model %T: %v", model, err)
			} else {
				log.Successf("Migration successful for model: %T", model)
			}
		}
	}
}
