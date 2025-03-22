package controller

import (
	"log"
	"strconv"
	"time"

	"gptbot/pkg/model"
	"gptbot/pkg/service"
	"gptbot/pkg/view"

	"github.com/gin-gonic/gin"
)

func EventView(c *gin.Context) (interface{}, error) {
	tzOffset := c.Query("tzoffset")
	filter := c.Query("filter")
	alert := c.Query("alert")
	return service.GetEventsTabView(tzOffset, filter, alert)
}

func EventCreate(c *gin.Context) (interface{}, error) {
	var newEvent view.NewEventFormView
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		return nil, err
	}

	event, err := service.BuildEventFromCreateView(&newEvent)
	if err != nil {
		return nil, err
	}

	if err := model.CreateEvent(event); err != nil {
		log.Printf("Failed to persist event: %v", err)
		return nil, err
	}

	event, err = model.ReadEvent(event.ID) // Reload with relations
	if err != nil {
		log.Printf("Failed to read event: %v", err)
	} else {
		log.Printf("New event registered: ID=%d Time=%s", 
			event.ID, event.GetTime().Format(time.RFC822))
		go service.PreprocessEvent(event)
	}

	return nil, onEventsChanged()
}

func EventUpdate(c *gin.Context) (interface{}, error) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 0, 0)
	if err != nil {
		return nil, err
	}

	var updEvent view.UpdateEventView
	if err := c.ShouldBindJSON(&updEvent); err != nil {
		return nil, err
	}

	event, err := service.BuildEventFromUpdateView(uint(id), &updEvent)
	if err != nil {
		return nil, err
	}

	if err := model.UpdateEvent(event); err != nil {
		return nil, err
	}

	return nil, onEventsChanged()
}

func EventDelete(c *gin.Context) (interface{}, error) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 0, 0)
	if err != nil {
		return nil, err
	}

	if err := model.DeleteEvent(uint(id)); err != nil {
		return nil, err
	}

	return nil, onEventsChanged()
}

func EventDeleteExpired(c *gin.Context) (interface{}, error) {
	events, err := model.ReadEvents()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	for _, ev := range *events {
		if ev.GetTime().Before(now) {
			model.DeleteEvent(ev.ID)
		}
	}

	return nil, onEventsChanged()
}

func onEventsChanged() error {
	if _, err := service.ScheduleEvents(); err != nil {
		return err
	}
	return nil
}