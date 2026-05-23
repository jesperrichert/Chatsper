package application

import (
	"log"

	"github.com/pterm/pterm"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	models "zip.jespersen.chatsper/internal/models/database"
	Log "zip.jespersen.chatsper/internal/utils"
	shared "zip.jespersen.chatsper/shared/database"
)

func NewDatabase(dbType shared.DatabaseType, connectionStr string) *gorm.DB {
	var database *gorm.DB

	Log.Debug("Connecting to database...")
	Log.Debug("Load database for type: " + pterm.Magenta(dbType))

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
	Log.Message("Running migrations for Database...")
	err := db.AutoMigrate(
		&models.UserEntity{},
		&models.TwitchUserEntity{},
	)
	if err != nil {
		Log.Error("Database Migrations failed to load for the current connection...", true)
	}
}
