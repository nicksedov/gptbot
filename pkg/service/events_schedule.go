package service

import (
	"gptbot/pkg/model"
	"gptbot/pkg/scheduler"
	ai "gptbot/pkg/localai"
	//ai "gptbot/pkg/gigachat"
)

func ScheduleEvents() (*[]model.SingleEvent, error) {
	events, dbErr := model.ReadEvents()
	if dbErr != nil {
		return nil, dbErr
	}
	scheduler.Schedule(events, ai.GetMessageByPrompt)
	return events, nil
}
