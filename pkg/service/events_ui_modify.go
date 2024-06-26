package service

import (
	"fmt"
	"strconv"
	"time"

	"gptbot/pkg/model"
	"gptbot/pkg/view"

	"gorm.io/datatypes"
)

func BuildEventFromCreateView(ev *view.NewEventFormView) (*model.SingleEvent, error) {
	var event *model.SingleEvent = &model.SingleEvent{}
	var params *[]model.SingleEventPromptParam = new([]model.SingleEventPromptParam)

	date, err := time.Parse(time.DateOnly, ev.Date)
	if err != nil {
		return nil, err
	}
	gormTime := new(datatypes.Time)
	err = gormTime.Scan(ev.Time)
	if err != nil {
		return nil, err
	}
	event.PromptID = ev.PromptID
	event.Date = datatypes.Date(date)
	event.Time = *gormTime
	event.TZOffset = ev.TZOffset
	event.TelegramChatID = ev.TelegramChatID

	params, err = appendParam(params, ev.ParamID0, ev.Param0)
	if err == nil {
		params, err = appendParam(params, ev.ParamID1, ev.Param1)
		if err == nil {
			params, err = appendParam(params, ev.ParamID2, ev.Param2)
		}
	}
	if err != nil {
		return nil, err
	}
	event.EventPromptParams = *params
	return event, nil
}

func appendParam(pp *[]model.SingleEventPromptParam, id, value string) (*[]model.SingleEventPromptParam, error) {
	if id != "" {
		id, uintErr := strconv.ParseUint(id, 0, 0)
		if uintErr != nil {
			return nil, uintErr
		}
		res := append(*pp, model.SingleEventPromptParam{Value: value, PromptParamID: uint(id)})
		return &res, nil
	}
	return pp, nil
}

func BuildEventFromUpdateView(id uint, ev *view.UpdateEventView) (*model.SingleEvent, error) {
	var event model.SingleEvent
	date, err := time.Parse(time.DateOnly, ev.Date)
	if err != nil {
		return nil, err
	}
	gormTime := new(datatypes.Time)
	err = gormTime.Scan(ev.Time)
	if err != nil {
		return nil, err
	}
	event.ID = id
	event.Date = datatypes.Date(date)
	event.Time = *gormTime
	event.TZOffset = ev.TZOffset
	event.TelegramChatID = ev.TelegramChatID
	return &event, nil
}

func DeleteEvent(id uint64) error {
	fmt.Printf("Deleting record with id=%d", id)
	return nil
}
