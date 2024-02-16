package model

import (
	"log"

	"github.com/nicksedov/gptbot/pkg/settings"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var sqliteDb *gorm.DB

func initDb() (*gorm.DB, error) {
	var err error
	if sqliteDb == nil {
		dbConfig := settings.GetSettings().DbConfig
		sqliteDb, err = gorm.Open(sqlite.Open(dbConfig.Path), &gorm.Config{})
	}
	return sqliteDb, err
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
