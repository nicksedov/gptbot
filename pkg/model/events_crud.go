package model

import (
	"fmt"

	"gorm.io/gorm"
)

func ReadEvents() (*[]SingleEvent, error) {
	findAllEvents := func(events *[]SingleEvent, db *gorm.DB) {
		db.Preload("EventPromptParams.PromptParam").Joins("Prompt").Joins("Chat").Find(events)
	}
	return read(findAllEvents)
}

func ReadEvent(eventId uint) (*SingleEvent, error) {
	findEvent := func(events *[]SingleEvent, db *gorm.DB) {
		db.Preload("EventPromptParams.PromptParam").Joins("Prompt").Joins("Chat").First(events, eventId)
	}
	result, err := read(findEvent)
	if err != nil {
		return nil, err
	} else if len(*result) > 0 {
		return &(*result)[0], nil
	} else {
		return nil, fmt.Errorf("event not found by id=%d", eventId)
	}
}

func CreateEvent(ev *SingleEvent) error {
	db, err := getDb()
	if err == nil {
		return db.Create(ev).Omit("Prompt", "Chat").Error
	}
	return err
}

func UpdateEvent(ev *SingleEvent) error {
	db, err := getDb()
	if err == nil {
		return db.Model(&ev).Omit("PromptID").Updates(ev).Error
	}
	return err
}

func DeleteEvent(id uint) error {
	db, err := getDb()
	if err == nil {
		return db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("\"singleEventId\" = ?", id).Delete(&SingleEventPromptParam{}).Error; err != nil {
				return err // return any error will rollback
			}
			if err := tx.Delete(&SingleEvent{ID: id}).Error; err != nil {
				return err
			}
			return nil
		})
	}
	return err
}
