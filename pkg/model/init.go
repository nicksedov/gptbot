package model

import (
	"log"

	"github.com/nicksedov/gptbot/pkg/settings"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	db.Preload("EventPromptParams.PromptParam").Preload("Chat").Preload(clause.Associations).Find(&events)
	return events, nil
}
