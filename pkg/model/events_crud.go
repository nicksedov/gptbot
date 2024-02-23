package model

import (
	"log"

	"gorm.io/gorm"
)

func ReadEvents() (*[]SingleEvent, error) {
	findAllEvents := func(events *[]SingleEvent, db *gorm.DB) {
		db.Preload("EventPromptParams.PromptParam").Joins("Prompt").Joins("Chat").Find(events)
	}
	return read(findAllEvents)
}

func CreateEvent(ev *SingleEvent) error {
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
