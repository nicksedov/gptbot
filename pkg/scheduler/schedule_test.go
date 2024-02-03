package scheduler

import (
	"testing"
	"time"

	"github.com/nicksedov/gptbot/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestSchedule(t *testing.T) {
	
	var handler *TestEventHandler = &TestEventHandler{}
	handler.triggerTime = time.Now()

	// Events count
	eventsToSchedule := 5
	eventsInterval := 100 * time.Millisecond
	// Expected that last event will be fired after given amount of time
	expectedCompletionTime := time.Duration(eventsToSchedule * int(eventsInterval))
	
	// Create events in the future
	var events []model.SingleEvent
	for i := 1; i <= eventsToSchedule; i++ {
		fireTime := time.Duration(i * int(eventsInterval))
		event := model.GetFutureTestEvent(fireTime)
		events = append(events, event)
	}

	// Run test
	Schedule(events, handler)
	
	// Evaluate real duration and compare with expected
	actualTime := handler.fireTime.Sub(handler.triggerTime)
	assert.GreaterOrEqual(t, actualTime, expectedCompletionTime)
	assert.LessOrEqual(t, actualTime, expectedCompletionTime + 100*time.Millisecond)
	assert.Equal(t, eventsToSchedule, handler.fireCount)
}

type TestEventHandler struct { 
	triggerTime time.Time
	fireTime time.Time
	fireCount int
}

func (h *TestEventHandler) handle(t *model.SingleEvent) {
	h.fireTime = time.Now()
	h.fireCount++
}
