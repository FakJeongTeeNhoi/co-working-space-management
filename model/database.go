package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var MainDB *gorm.DB

func InitDB() {
	var err error

	MainDB, err = gorm.Open(postgres.Open(
		"host="+os.Getenv("DB_HOST")+
			" user="+os.Getenv("DB_USER")+
			" password="+os.Getenv("DB_PASSWORD")+
			" dbname="+os.Getenv("DB_NAME")+
			" port="+os.Getenv("DB_PORT")+
			"TimeZone=Asia/Bangkok"), &gorm.Config{})

	if err != nil {
		log.Fatalln("Unable to connect to database: ", err)
	} else {
		log.Println("Connected to database")
	}

	MigrateDB()
}

func MigrateDB() {
	// TODO: Add models here
	Models := []interface{}{}

	err := MainDB.AutoMigrate(Models...)
	if err != nil {
		log.Fatalln("Unable to migrate database: ", err)
	} else {
		log.Println("Migrated database")
	}
}
