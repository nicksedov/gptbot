package service

import (
	"sort"
	"strconv"
	"time"

	"gptbot/pkg/model"
	"gptbot/pkg/view"
)

func GetEventsTabView(offsetParam string, filter string, alert string) (*view.EventsTabView, error) {
	events, dbErr := model.ReadEvents()
	if dbErr != nil {
		return nil, dbErr
	}
	chats, dbErr := model.GetAll[model.TelegramChat]()
	if dbErr != nil {
		return nil, dbErr
	}
	chatItems := getAsDropdown(chats)

	eventViewMap := make(map[uint]view.EventView, len(*events))
	eventOrderMap := make(map[uint]time.Time, len(*events))
	for _, ev := range *events {
		evTime := ev.GetTime()
		if !timeFilterPredicate(evTime, filter) {
			continue
		}
		tzOffset, differentTimeZone := parseOffsetParam(offsetParam, ev.TZOffset)
		if differentTimeZone {
			evTime = evTime.In(time.FixedZone("", -tzOffset*60))
		}
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
	promptParamViews := make([]view.EventPromptParam, len(*promptParams))
	for i, promptParamItem := range *promptParams {
		promptParamViews[i] = view.EventPromptParam{ID: promptParamItem.ID, PromptID: promptParamItem.PromptID, Title: promptParamItem.Title}
	}

	return &view.EventsTabView{
		Alert: alert,
		EventViews: eventViews, 
		Prompts: promptItems, 
		PromptParams: promptParamViews, 
		Chats: chatItems,
	}, nil
}

func timeFilterPredicate(evTime time.Time, filter string) bool {
	currentTime := time.Now()
	switch filter {
	case "past":
		return evTime.Before(currentTime)
	case "future":
		return evTime.After(currentTime)
	case "today":
		return isToday(evTime)
	case "tomorrow":
		return isToday(evTime.Add(-24 * time.Duration(time.Hour)))
	default:
		return true
	}
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

func getAsDropdown[T model.IDValue](idValueList *[]T) []view.DropdownItem {
	listItems := make([]view.DropdownItem, len(*idValueList))
	for i, idValue := range *idValueList {
		listItems[i] = view.DropdownItem{ID: idValue.GetId(), Value: idValue.GetValue()}
	}
	return listItems
}

func parseOffsetParam(offsetParam string, defaultVal int) (int, bool) {
	if offsetParam != "" {
		intOffset, err := strconv.Atoi(offsetParam)
		if (err == nil) && (intOffset != defaultVal) {
			return intOffset, true
		}
	}
	return defaultVal, false
}

func isToday(evTime time.Time) bool {
	now := time.Now()
	localTime := evTime.In(time.Local)
	return (localTime.Year() == now.Year()) && (localTime.Month() == now.Month()) && (localTime.Day() == now.Day())
}
