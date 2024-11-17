package model

import (
	"errors"

	"gorm.io/gorm"
)

func CreatePrebuiltMessage(msg *SingleEventPrebuiltMessage) error {
	db, err := getDb()
	if err == nil {
		return db.Create(msg).Error
	}
	return err
}

func UpdatePrebuiltMessage(msg *SingleEventPrebuiltMessage) error {
	db, err := getDb()
	if err == nil {
		return db.Model(&msg).Updates(msg).Error
	}
	return err
}

func DeletePrebuiltMessages(eventId uint) error {
	db, err := getDb()
	if err == nil {
		return db.Where("\"singleEventId\" = ?", eventId).Delete(&SingleEventPrebuiltMessage{}).Error
	}
	return err
}

func ReadPrebuiltMessageByEventId(eventId uint) (*SingleEventPrebuiltMessage, error) {
	findAllMessages := func(msgList *[]SingleEventPrebuiltMessage, db *gorm.DB)  *gorm.DB {
		return db.Where(&SingleEventPrebuiltMessage{EventID: eventId}).Find(msgList)
	}
	allMessages, err := readMany(findAllMessages)
	if err == nil {
		var pending, failed *SingleEventPrebuiltMessage
		for _, msg := range *allMessages {
			if msg.Status == Created {
				return &msg, nil
			} else if msg.Status == Pending {
				pending = &msg
			} else if msg.Status == Failed {
				failed = &msg
			}
		}
		if pending != nil {
			return pending, nil
		} else if failed != nil {
			return failed, nil
		}
		return nil, errors.New("prebuilt messages not found for given event")
	}
	return nil, err
}
