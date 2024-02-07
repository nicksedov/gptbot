package model

import (
	"log"

	"github.com/nicksedov/gptbot/pkg/settings"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var events []SingleEvent

func GetEvents() ([]SingleEvent, error) {
	dbConfig := settings.GetSettings().DbConfig
	sqliteDb, err := gorm.Open(sqlite.Open(dbConfig.Path), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}
	sqliteDb.Preload("EventPromptParams.PromptParam").Preload("Chat").Preload(clause.Associations).Find(&events)
	return events, nil
}
