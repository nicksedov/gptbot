package scheduler

import (
	"time"

	"gptbot/pkg/model"

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

func Schedule(events *[]model.SingleEvent, getMessageByPrompt func(e *model.SingleEvent) (string, error)) {

	scheduler := initScheduler()

	for _, event := range *events {
		fireTime := event.GetTime()
		duration := time.Until(fireTime)
		if duration > 0 {
			task := tasks.Task{
				Interval: duration,
				RunOnce:  true,
				TaskFunc: func() error { return handle(&event, getMessageByPrompt) },
				ErrFunc:  onError,
			}
			scheduler.Add(&task)
		}
	}
}
