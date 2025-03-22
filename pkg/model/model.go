package model

import (
	"time"

	"gorm.io/datatypes"
)

type Prompt struct {
	ID      uint `gorm:"unique;primaryKey;autoIncrement"`
	Title   string
	Prompt  string
	AltText string `gorm:"column:altText"`
	PromptParams []PromptParam  `gorm:"foreignKey:PromptID"` // Добавляем связь
}

func (p Prompt) GetId() uint {
	return p.ID
}
func (p Prompt) GetValue() string {
	return p.Title
}

type PromptParam struct {
	ID       uint `gorm:"unique;primaryKey;autoIncrement"`
	Tag      string
	Title    string
	PromptID uint `gorm:"column:promptId;index"`
}

type TelegramChat struct {
	ID       uint   `gorm:"unique;primaryKey;autoIncrement"`
	ChatID   int64  `gorm:"column:chatId"`
	ChatName string `gorm:"column:chatName"`
}

func (tc TelegramChat) GetId() uint {
	return tc.ID
}
func (tc TelegramChat) GetValue() string {
	return tc.ChatName
}

type SingleEvent struct {
	ID                uint `gorm:"unique;primaryKey;autoIncrement"`
	Date              datatypes.Date
	Time              datatypes.Time
	TZOffset          int                      `gorm:"column:tzOffset"`
	TelegramChatID    uint                     `gorm:"column:telegramChatId"`
	PromptID          uint                     `gorm:"column:promptId"`
	Prompt            Prompt                   `gorm:"foreignKey:PromptID"`
	EventPromptParams []SingleEventPromptParam `gorm:"foreignKey:EventID"`
	Chat              TelegramChat             `gorm:"foreignKey:TelegramChatID"`
}

type SingleEventPromptParam struct {
	ID            uint `gorm:"unique;primaryKey;autoIncrement"`
	Value         string
	EventID       uint        `gorm:"column:singleEventId"`
	PromptParamID uint        `gorm:"column:promptParamId"`
	PromptParam   PromptParam `gorm:"foreignKey:PromptParamID"`
}

type MessageBuildStatus int

const (
	Pending MessageBuildStatus = iota
	Created
	Failed
)

type SingleEventPrebuiltMessage struct {
	ID          uint               `gorm:"unique;primaryKey;autoIncrement"`
	EventID     uint               `gorm:"column:singleEventId"`
	Status      MessageBuildStatus `gorm:"column:status"`
	Message     string             `gorm:"column:message"`
	RequestedAt time.Time          `gorm:"column:requestedAt"`
	BuiltAt     time.Time          `gorm:"column:builtAt"`
}