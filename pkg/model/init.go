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

func Get[T any](finder func(items *[]T, db *gorm.DB)) (*[]T, error) {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}
	items := new([]T)
	finder(items, db)
	return items, nil
} 

func GetEvents() (*[]SingleEvent, error) {
	return Get(func(events *[]SingleEvent, db *gorm.DB) {
		db.Preload("EventPromptParams.PromptParam").Joins("Prompt").Joins("Chat").Find(events)
	})	
}

func GetAll[T any]() (*[]T, error) {
	return Get(func(items *[]T, db *gorm.DB) {
		db.Find(items)
	})
}

func AddEvent(ev *SingleEvent) error {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return err
	}
	tx := db.Create(ev).Omit("Prompt", "Chat")
	tx.Commit()
	return nil
}

func UpdateEvent(ev *SingleEvent) error {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return err
	}
	tx := db.Model(&ev).Omit("PromptID").Updates(ev)
	tx.Commit()
	return nil
}

func DeleteEvent(id uint) error {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return err
	}
	db.Where("\"singleEventId\" = ?", id).Delete(&SingleEventPromptParam{})
	db.Delete(&SingleEvent{ID: id})
	return nil
}