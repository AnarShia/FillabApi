package database

import (
	"fmt"
	"log"
	"os"

	"github.com/AnarShia/FillabApi/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var Db Dbinstance

func ConnectDbSqlite() {

	fmt.Println("Connecting to database...")
	db, error := gorm.Open(sqlite.Open("Fillab_AnarShia.db"), &gorm.Config{})
	if error != nil {
		log.Fatal("Failed to connect to database!. \n", error)
		os.Exit(2)
	}

	log.Println("Connected to database!")
	db.Logger = db.Logger.LogMode(logger.Info)

	log.Println("Migrating database...")
	db.AutoMigrate(&models.User{})

	Db = Dbinstance{
		Db: db,
	}

}
