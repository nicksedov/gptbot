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

type PromptParam struct {
	ID       uint `gorm:"primaryKey"`
	Tag      string
	Title    string
	PromptID uint `gorm:"column:promptId"`
}

func (PromptParam) TableName() string {
	return "prompt_params"
}

type SingleEvent struct {
	ID                uint `gorm:"primaryKey"`
	Date              datatypes.Date
	Time              datatypes.Time
	TZOffset          int                      `gorm:"column:tzOffset"`
	PromptID          uint                     `gorm:"column:promptId"`
	Prompt            Prompt                   `gorm:"foreignKey:PromptID"`
	EventPromptParams []SingleEventPromptParam `gorm:"foreignKey:EventID"`
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
