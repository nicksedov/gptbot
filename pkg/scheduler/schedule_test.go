package scheduler

import (
	"gptbot/pkg/model"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSchedule(t *testing.T) {

	var handler *TestEventHandler = &TestEventHandler{}
	handler.triggerTime = time.Now()

	// Events count
	eventsToSchedule := 5
	eventsInterval := 100 * time.Millisecond
	// Expected that last event will be fired after given amount of time
	expectedCompletionTime := time.Duration(int64(eventsInterval) * int64(eventsToSchedule))

	// Create events in the future
	var events []model.SingleEvent
	for i := 1; i <= eventsToSchedule; i++ {
		fireTime := time.Duration(i * int(eventsInterval))
		event := model.GetFutureTestEvent(uint(i), fireTime)
		events = append(events, event)
	}

	// Run test
	Schedule(&events, handler.testFunc)

	overlapTime := 1000 * time.Millisecond
	// Scheduled events execute asynchronously in goroutines; need to wait enough time until they complete
	time.Sleep(expectedCompletionTime + overlapTime)
	log.Printf("Verifying that %d background events successfully fired at this moment\n", eventsToSchedule)
	// Evaluate real duration and compare with expected
	actualTime := handler.fireTime.Sub(handler.triggerTime)
	assert.GreaterOrEqual(t, actualTime, expectedCompletionTime)
	assert.LessOrEqual(t, actualTime, expectedCompletionTime+overlapTime)
	assert.Equal(t, eventsToSchedule, handler.fireCount)
}

type TestEventHandler struct {
	triggerTime time.Time
	fireTime    time.Time
	fireCount   int
}

func (h *TestEventHandler) testFunc(t *model.SingleEvent) error {
	h.fireTime = time.Now()
	h.fireCount++
	log.Printf("Firing background event with ID=%d\n", t.ID)
	return nil
}
