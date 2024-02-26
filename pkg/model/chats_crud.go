package model

import "log"

func CreateChat(c *TelegramChat) error {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return err
	}
	tx := db.Create(c)
	tx.Commit()
	return nil
}

func DeleteChat(id uint) error {
	db, err := initDb()
	if err != nil {
		log.Fatal("failed to connect database")
		return err
	}
	db.Delete(&TelegramChat{ID: id})
	return nil
}
