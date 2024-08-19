package service

import (
	"gptbot/pkg/model"
	"gptbot/pkg/scheduler"
)

func ScheduleEvents() (*[]model.SingleEvent, error) {
	events, dbErr := model.ReadEvents()
	if dbErr != nil {
		return nil, dbErr
	}
	//var h scheduler.EventHandler = &scheduler.GigaChatEventHandler{}
	var h scheduler.EventHandler = &scheduler.LocalAIEventHandler{}
	scheduler.Schedule(events, h)
	return events, nil
}
