package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nicksedov/gptbot/pkg/model"
	"github.com/nicksedov/gptbot/pkg/service"
	"github.com/nicksedov/gptbot/pkg/view"
)

func EventView(c *gin.Context) {
	var tzOffset string = c.Query("tzoffset")
	eventsTab, err := service.GetEventsTabView(tzOffset)
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
		err = model.AddEvent(event)
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

func onEventsChanged(c *gin.Context, err error) {
	if err == nil {
		service.LoadAndScheduleEvents()
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}
