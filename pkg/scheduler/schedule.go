package scheduler

import (
	"time"

	"github.com/madflojo/tasks"
	"github.com/nicksedov/gptbot/pkg/model"
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

func Schedule(events []model.SingleEvent, h EventHandler) {
	
	scheduler := initScheduler()

	for _, event := range events {
		fireTime := event.GetTime()
		duration := time.Until(fireTime)
		if duration > 0 {
			task := tasks.Task {
				Interval: duration,
				RunOnce: true,
				TaskFunc: func() error { return h.handle(&event) },
				ErrFunc: h.onError,
			}
			scheduler.Add(&task)
		}
	}
}