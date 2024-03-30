package controller

import (
	"net/http"
	"strconv"
	"time"

	"gptbot/pkg/model"
	"gptbot/pkg/service"
	"gptbot/pkg/view"

	"github.com/gin-gonic/gin"
)

func EventView(c *gin.Context) {
	var tzOffset string = c.Query("tzoffset")
	var filter string = c.Query("filter")
	eventsTab, err := service.GetEventsTabView(tzOffset, filter)
	if err == nil {
		c.JSON(http.StatusOK, eventsTab)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}

func EventCreate(c *gin.Context) {
	var newEvent view.NewEventFormView
	c.ShouldBindJSON(&newEvent)
	event, err := service.BuildEventFromCreateView(&newEvent)
	if err == nil {
		err = model.CreateEvent(event)
	}
	onEventsChanged(c, err)
}

func EventUpdate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 0, 0)
	if err == nil {
		var updEvent view.UpdateEventView
		c.ShouldBindJSON(&updEvent)
		event, updErr := service.BuildEventFromUpdateView(uint(id), &updEvent)
		if updErr == nil {
			updErr = model.UpdateEvent(event)
		}
		err = updErr
	}
	onEventsChanged(c, err)
}

func EventDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 0, 0)
	if err == nil {
		err = model.DeleteEvent(uint(id))
	}
	onEventsChanged(c, err)
}

func EventDeleteExpired(c *gin.Context) {
	events, err := model.ReadEvents()
	if err == nil {
		now := time.Now()
		for _, ev := range *events {
			if ev.GetTime().Before(now) {
				model.DeleteEvent(ev.ID)
			}
		}
	}
	onEventsChanged(c, err)
}

func onEventsChanged(c *gin.Context, err error) {
	if err == nil {
		service.ScheduleEvents()
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}
