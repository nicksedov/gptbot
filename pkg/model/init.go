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
			dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.User, dbConfig.Password, dbConfig.SSLMode)
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

func read[T any](selector func(items *[]T, db *gorm.DB)) (*[]T, error) {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return nil, err
	}
	items := new([]T)
	selector(items, db)
	return items, nil
}

func GetAll[T any]() (*[]T, error) {
	selectAll := func(items *[]T, db *gorm.DB) {
		db.Find(items)
	}
	return read(selectAll)
}
