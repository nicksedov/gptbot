package service

import (
	"gptbot/pkg/model"
	"gptbot/pkg/scheduler"
)

func ScheduleEvents() (*[]model.SingleEvent, error) {
	events, dbErr := model.ReadEvents()
	if dbErr == nil {
		scheduler.Schedule(events, scheduler.HandleEvent)
		return events, nil
	}
	return nil, dbErr
}
