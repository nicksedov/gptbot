package scheduler

import (
	"sync"
	"time"

	"github.com/nicksedov/gptbot/pkg/model"
)

func Schedule(events []model.SingleEvent, h EventHandler) {
	var schedWaitGroup sync.WaitGroup
	for _, event := range events {
		chatId := 1 // TODO evaluate chat ID for event
		if chatId != 0 {
			fireTime := event.GetTime()
			duration := time.Until(fireTime)
			if duration > 0 {
				schedWaitGroup.Add(1)
				go func(ev *model.SingleEvent) {
					defer schedWaitGroup.Done()
					time.Sleep(duration)
					h.handle(ev)
				}(&event)
			}
		}
	}
	schedWaitGroup.Wait()
}
