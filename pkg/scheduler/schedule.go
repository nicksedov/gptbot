package scheduler

import (
	"gptbot/pkg/model"
	"gptbot/pkg/telegram"
	"time"

	"github.com/madflojo/tasks"
)

var scheduler *tasks.Scheduler

func initScheduler() *tasks.Scheduler {
	if scheduler == nil {
		scheduler = tasks.New()
	} else {
		scheduler.Stop()
	}
	return scheduler
}

func Schedule(events *[]model.SingleEvent, handler func(event *model.SingleEvent) error) {

	scheduler := initScheduler()

	for _, event := range *events {
		fireTime := event.GetTime()
		duration := time.Until(fireTime)
		if duration > 0 {
			task := tasks.Task{
				Interval: duration,
				RunOnce:  true,
				TaskFunc: func() error { return handler(&event) },
				ErrFunc:  telegram.ErrorReporter,
			}
			scheduler.Add(&task)
		}
	}
}
