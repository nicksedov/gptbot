package model

import (
	"gorm.io/datatypes"
)

type Prompt struct {
	ID      uint `gorm:"primaryKey"`
	Title   string
	Prompt  string
	AltText string `gorm:"column:altText"`
}
func (p Prompt) GetId() uint {
	return p.ID
}
func (p Prompt) GetValue() string {
	return p.Title
}

type PromptParam struct {
	ID       uint `gorm:"primaryKey"`
	Tag      string
	Title    string
	PromptID uint `gorm:"column:promptId"`
}
func (PromptParam) TableName() string {
	return "prompt_params"
}

type TelegramChat struct {
	ID       uint   `gorm:"primaryKey"`
	ChatID   int64  `gorm:"column:chatId"`
	ChatName string `gorm:"column:chatName"`
}

func (TelegramChat) TableName() string {
	return "telegram_chats"
}
func (tc TelegramChat) GetId() uint {
	return tc.ID
}
func (tc TelegramChat) GetValue() string {
	return tc.ChatName
}

type SingleEvent struct {
	ID                uint `gorm:"primaryKey"`
	Date              datatypes.Date
	Time              datatypes.Time
	TZOffset          int                      `gorm:"column:tzOffset"`
	TelegramChatID    uint                     `gorm:"column:telegramChatId"`
	PromptID          uint                     `gorm:"column:promptId"`
	Prompt            Prompt                   `gorm:"foreignKey:PromptID"`
	EventPromptParams []SingleEventPromptParam `gorm:"foreignKey:EventID"`
	Chat              TelegramChat             `gorm:"foreignKey:TelegramChatID"`
}

func (SingleEvent) TableName() string {
	return "single_events"
}

type SingleEventPromptParam struct {
	ID            uint `gorm:"primaryKey"`
	Value         string
	EventID       uint        `gorm:"column:singleEventId"`
	PromptParamID uint        `gorm:"column:promptParamId"`
	PromptParam   PromptParam `gorm:"foreignKey:PromptParamID"`
}

func (SingleEventPromptParam) TableName() string {
	return "single_event_prompt_params"
}
