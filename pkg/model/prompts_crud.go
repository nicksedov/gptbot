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

func GetPromptParamsByPromptID(promptID uint) ([]PromptParam, error) {
	db, err := getDb()
	if err != nil {
		return nil, err
	}
	
	var params []PromptParam
	err = db.Where("prompt_id = ?", promptID).Find(&params).Error
	return params, err
}

func UpdatePromptParam(param *PromptParam) error {
	db, err := getDb()
	if err == nil {
		return db.Save(param).Error
	}
	return err
}

func CreatePromptParam(param *PromptParam) error {
	db, err := getDb()
	if err == nil {
		return db.Create(param).Error
	}
	return err
}

func DeletePromptParam(id uint) error {
	db, err := getDb()
	if err == nil {
		return db.Delete(&PromptParam{ID: id}).Error
	}
	return err
}

func GetPromptWithParams(id uint) (*Prompt, error) {
	db, err := getDb()
	if err != nil {
		return nil, err
	}
	var prompt Prompt
	err = db.Preload("PromptParams").
		First(&prompt, id).
		Error

	return &prompt, err
}

func UpdatePrompt(p *Prompt) error {
	db, err := getDb()
	if err == nil {
		return db.Save(p).Error
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
