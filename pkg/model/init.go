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
	db.Preload("EventPromptParams.PromptParam").Joins("Chat").Preload(clause.Associations).Find(&events)
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

func GetChats() ([]TelegramChat, error) {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}
	var chats []TelegramChat
	db.Find(&chats)
	return chats, nil
}

func GetPrompts() ([]Prompt, error) {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}
	var prompts []Prompt
	db.Find(&prompts)
	return prompts, nil
}

func GetPromptParams() ([]PromptParam, error) {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}
	var promptParams []PromptParam
	db.Find(&promptParams)
	return promptParams, nil
}
