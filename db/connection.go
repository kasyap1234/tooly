package database

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection and migrates models
func ConnectDB(dsn string, models ...interface{}) {
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    // Migrate the schema for all provided models
    if err := DB.AutoMigrate(models...); err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }
}