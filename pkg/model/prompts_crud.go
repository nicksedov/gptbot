package model

import "gorm.io/gorm"

func CreatePrompt(p *Prompt) error {
	db, err := getDb()
	if err == nil {
		return db.Create(p).Error
	}
	return err
}

func CreatePromptParams(p []PromptParam) error {
	db, err := getDb()
	if err == nil {
		return db.CreateInBatches(p, 10).Error
	}
	return err
}

func DeletePrompt(id uint) error {
	db, err := getDb()
	if err == nil {
		return db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("\"promptId\" = ?", id).Delete(&PromptParam{}).Error; err != nil {
				return err // return any error will rollback
			}
			if err := tx.Delete(&Prompt{ID: id}).Error; err != nil {
				return err // return any error will rollback
			}
			return nil
		})
	}
	return err
}
