package database

import (
	"log"
	"os"

	"github.com/aalekh12/Blog-Post-Api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeSettings() models.Settings {
	DB_HOST := os.Getenv("DB_HOST")
	DB_TYPE := os.Getenv("DB_TYPE")
	SQLITE_DB_NAME := os.Getenv("SQLITE_DB_NAME")

	switch {
	case DB_HOST == "":
		log.Println("Environment variable MYSQL_DB_HOST not set.")
		os.Exit(1)
	case SQLITE_DB_NAME == "":
		log.Println("Environment variable SQLITE_DB_NAME not set.")
		os.Exit(1)
	case DB_TYPE == "":
		log.Println("Environment variable DB_TYPE not set.")
		os.Exit(1)
	default:
		log.Println("All environment variables are set correctly.")
	}

	settings := models.Settings{
		DB_HOST:        DB_HOST,
		DB_TYPE:        DB_TYPE,
		SQLITE_DB_NAME: SQLITE_DB_NAME,
	}

	return settings
}

func ConnectDataBase() {
	log.Println("(db): Initializing settings")
	settings := InitializeSettings()

	log.Printf("(db): Initializing DB for %v\n", settings.DB_TYPE)
	switch settings.DB_TYPE {
	case "sqlite":
		db, err := gorm.Open(sqlite.Open(settings.SQLITE_DB_NAME), &gorm.Config{})
		log.Println("(db) SQLITE open completed")
		if err != nil {
			log.Printf("(db) SQLITE open error (%s)", err)
			return
		}
		log.Println("(db) Migrating DB ")
		db.AutoMigrate(
			&models.Blog{},
		)

		DB = db
	default:
		panic("invalid database provided")
	}

	log.Println("(db): DB init complete.")
}
