package application

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"zip.jespersen.chatsper/Chatsper.Backend/internal/models"
	shared "zip.jespersen.chatsper/Chatsper.Backend/shared/database"
)

func NewDatabase(dbType shared.DatabaseType, connectionStr string) *gorm.DB {
	var database *gorm.DB

	switch dbType {
	case shared.SqliteDataBase:
		{
			db, err := gorm.Open(sqlite.Open("chatsper.db"), &gorm.Config{})
			if err != nil {
				log.Fatalf("failed to connect to database: %v", err)
			}
			database = db
		}
		break
	case shared.PostgresDatabase:
		{
			db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
			if err != nil {
				log.Fatalf("failed to connect to database: %v", err)
			}
			database = db
		}
		break
	}

	migrations(database)
	return database
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.ExampleEntity{})
	if err != nil {
		return
	}
}
