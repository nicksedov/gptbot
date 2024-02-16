package service

import (
	"sort"
	"time"

	"github.com/nicksedov/gptbot/pkg/model"
	"github.com/nicksedov/gptbot/pkg/view"
)

func GetEventsTabView() (*view.EventsTabView, error) {
	events, dbErr := model.GetEvents()
	if dbErr != nil {
		return nil, dbErr
	}
	chats, dbErr := model.GetAll[model.TelegramChat]()
	if dbErr != nil {
		return nil, dbErr
	}
	chatItems := getAsDropdown(chats)

	eventViewMap := make(map[uint]view.EventView, len(events))
	eventOrderMap := make(map[uint]time.Time, len(events))
	for _, ev := range events {
		evTime := ev.GetTime()
		tzOffset := ev.TZOffset
		prompt, err := ev.GetResolvedPrompt()
		if err != nil {
			continue
		}
		eventViewMap[ev.ID] = view.EventView{
			ID:             ev.ID,
			Date:           evTime.Format(time.DateOnly),
			Time:           evTime.Format(time.TimeOnly),
			TZOffset:       tzOffset,
			PromptTitle:    ev.Prompt.Title,
			Prompt:         prompt,
			TelegramChatID: ev.TelegramChatID,
		}
		eventOrderMap[ev.ID] = evTime
	}

	orderedKeys := orderIDsByTime(eventOrderMap)
	eventViews := make([]view.EventView, len(orderedKeys))
	for i, key := range orderedKeys {
		eventViews[i] = eventViewMap[key]
	}

	prompts, dbErr := model.GetAll[model.Prompt]()
	if dbErr != nil {
		return nil, dbErr
	}
	promptItems := getAsDropdown(prompts)

	promptParams, dbErr := model.GetAll[model.PromptParam]()
	if dbErr != nil {
		return nil, dbErr
	}
	promptParamViews := make([]view.PromptParamView, len(promptParams))
	for i, promptParamItem := range promptParams {
		promptParamViews[i] = view.PromptParamView{ID: promptParamItem.ID, PromptID: promptParamItem.PromptID, Title: promptParamItem.Title}
	} 

	return &view.EventsTabView{EventViews: eventViews, Prompts: promptItems, PromptParams: promptParamViews, Chats: chatItems}, nil
}

func orderIDsByTime(idByTimeMap map[uint]time.Time) []uint {
	keys := make([]uint, 0, len(idByTimeMap))
	for key := range idByTimeMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return idByTimeMap[keys[i]].Before(idByTimeMap[keys[j]])
	})
	return keys
}

func getAsDropdown[T model.IDValue](idValueList []T) []view.DropdownItem {
	listItems := make([]view.DropdownItem, len(idValueList))
	for i, idValue := range idValueList {
		listItems[i] = view.DropdownItem{ID: idValue.GetId(), Value: idValue.GetValue()}
	}
	return listItems
}
