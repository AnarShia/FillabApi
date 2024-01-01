package database

import (
	"fmt"
	"log"
	"os"

	"github.com/AnarShia/FillabApi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var Db Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if error != nil {
		log.Fatal("Failed to connect to database!. \n", error)
		os.Exit(2)
	}
	log.Println("Connected to database!")
	db.Logger = db.Logger.LogMode(logger.Info)

	log.Println("Migrating database...")
	db.AutoMigrate(&models.Fact{})

	Db = Dbinstance{
		Db: db,
	}
}
