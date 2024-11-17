package model

func CreateChat(c *TelegramChat) error {
	db, err := getDb()
	if err == nil {
		return db.Create(c).Error
	}
	return err
}

func DeleteChat(id uint) error {
	db, err := getDb()
	if err == nil {
		return db.Delete(&TelegramChat{ID: id}).Error
	}
	return err
}

func GetChat(id uint) (*TelegramChat, error) {
	db, err := getDb()
	if err != nil {
		return nil, err
	}
	result := &TelegramChat{}
	tx := db.First(result, id)
	if tx.Error == nil {
		return result, nil
	} else {
		return nil, tx.Error
	}
}
