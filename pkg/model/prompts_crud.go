package model

import "log"

func CreatePrompt(p *Prompt) error {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return err
	}
	tx := db.Create(p)
	tx.Commit()
	return nil
}

func CreatePromptParams(p []PromptParam) error {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return err
	}
	tx := db.CreateInBatches(p, 10)
	tx.Commit()
	return nil
}

func DeletePrompt(id uint) error {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return err
	}
	db.Where("\"promptId\" = ?", id).Delete(&PromptParam{})
	db.Delete(&Prompt{ID: id})
	return nil
}
