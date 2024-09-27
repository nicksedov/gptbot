package model

import (
	"fmt"
	"log"

	"gptbot/pkg/settings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func initDb() (*gorm.DB, error) {
	dbConfig := settings.GetSettings().DbConfig
	dsnFormat := "host=%s port=%d dbname=%s user=%s password=%s sslmode=%s"
	dsn := fmt.Sprintf(dsnFormat,
		dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.User, dbConfig.Password, dbConfig.SSLMode)
	log.Printf("Opening database connection: postgres://%s:%d/%s?sslMode=%s\n", dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.SSLMode)	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		db.AutoMigrate(
			&Prompt{},
			&PromptParam{},
			&SingleEvent{},
			&SingleEventPromptParam{},
			&TelegramChat{},
			&SingleEventPrebuiltMessage{},
		)
	}
	return db, err
}

func getDb() (*gorm.DB, error) {
	var err error
	if database == nil {
		dbConfig := settings.GetSettings().DbConfig
		database, err = initDb()
		if err != nil {
			log.Printf("Failed to connect database postgres://%s:%d/%s?sslMode=%s\n", dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.SSLMode)
		} else {
			tx := database.Exec("select 1;")
			if tx.Error == nil {
				log.Println("Database connection opened successfully")
			} else {
				log.Printf("Failed to access database postgres://%s:%d/%s?sslMode=%s\n", dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.SSLMode)
			}
			
		}
	}
	return database, err
}

func read[T any](selector func(items *[]T, db *gorm.DB)) (*[]T, error) {
	db, err := getDb()
	if err == nil {
		items := new([]T)
		selector(items, db)
		return items, nil
	}
	return nil, err
}

func GetAll[T any]() (*[]T, error) {
	selectAll := func(items *[]T, db *gorm.DB) {
		db.Order("id").Find(items)
	}
	return read(selectAll)
}