package model

import (
	"time"

	"gorm.io/datatypes"
)

var (
	testDate, _ = time.Parse(time.DateOnly, "2024-01-02")
	hour        = 10
	min         = 30
	sec         = 00
	ns          = 00
	tzOffset    = -180 // Timezone offset for "Europe/Moscow" location, as returned by JavaScript method getTimezoneOffset() 
)

var event = SingleEvent{
	ID:       0,
	Date:     datatypes.Date(testDate),
	Time:     datatypes.NewTime(hour, min, sec, ns),
	TZOffset: tzOffset,
	PromptID: 0,
	Prompt: Prompt{
		ID:      0,
		Title:   "Промпт для тестирования",
		Prompt:  "В этой строке проверяется корректность подстановки значений параметров '${first}' и '${second}'",
		AltText: "В этой строке проверяется альтернативный текст",
	},
	EventPromptParams: []SingleEventPromptParam{
		{
			ID:            0,
			Value:         "Значение первого параметра",
			EventID:       0,
			PromptParamID: 0,
			PromptParam: PromptParam{
				ID:       0,
				Tag:      "first",
				Title:    "Первый параметр",
				PromptID: 0,
			},
		},
		{
			ID:            1,
			Value:         "Значение второго параметра",
			EventID:       0,
			PromptParamID: 0,
			PromptParam: PromptParam{
				ID:       1,
				Tag:      "second",
				Title:    "Второй параметр",
				PromptID: 0,
			},
		},
	},
}

func GetDefaultTestEvent() SingleEvent {
	return event
}

func GetFutureTestEvent(d time.Duration) SingleEvent {
	futureEvent := event
	future := time.Now().Add(d)
	futureDate := future.Truncate(24 * time.Hour)
	futureEvent.Date = datatypes.Date(futureDate)
	futureEvent.Time = datatypes.NewTime(future.Hour(), future.Minute(), future.Second(), future.Nanosecond())
	return futureEvent
}

func GetTestEventAtMoment(yr int, mon time.Month, dt, h, m, s, ns int) SingleEvent {
	eventAtMoment := event
	eventAtMoment.Date = datatypes.Date(time.Date(yr, mon, dt, 0, 0, 0, 0, time.Local))
	eventAtMoment.Time = datatypes.NewTime(h, m, s, ns)
	return eventAtMoment
}
