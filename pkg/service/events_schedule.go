package service

import (
	"github.com/nicksedov/gptbot/pkg/model"
	"github.com/nicksedov/gptbot/pkg/scheduler"
)

func ScheduleEvents() (*[]model.SingleEvent, error) {
	events, dbErr := model.ReadEvents()
	if dbErr != nil {
		return nil, dbErr
	}
	var h *scheduler.GptChatEventHandler = &scheduler.GptChatEventHandler{}
	scheduler.Schedule(events, h)
	return events, nil
}
