package model

import (
	"fmt"
	"log"

	"github.com/nicksedov/gptbot/pkg/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDb() (*gorm.DB, error) {
	var err error
	if db == nil {
		dbConfig := settings.GetSettings().DbConfig
		dsnFormat := "host=%s port=%d dbname=%s user=%s password=%s sslmode=%s"
		dsn := fmt.Sprintf(dsnFormat, 
			dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.User, dbConfig.Password, dbConfig.SslMode) 
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		db.AutoMigrate(
			&Prompt{}, 
			&PromptParam{}, 
			&SingleEvent{}, 
			&SingleEventPromptParam{},
			&TelegramChat{})
	}
	return db, err
}

func GetEvents() ([]SingleEvent, error) {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}
	var events []SingleEvent
	db.Preload("EventPromptParams.PromptParam").Joins("Prompt").Joins("Chat").Find(&events)
	return events, nil
}

func GetAll[T any]() ([]T, error) {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}
	var items []T
	db.Find(&items)
	return items, nil
}

func AddEvent(ev SingleEvent) error {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return err
	}
	tx := db.Create(&ev).Omit("Prompt", "Chat")
	tx.Commit()
	return nil
}
