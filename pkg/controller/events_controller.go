package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nicksedov/gptbot/pkg/model"
	"github.com/nicksedov/gptbot/pkg/service"
	"github.com/nicksedov/gptbot/pkg/view"
)

func EventList(c *gin.Context) {
	events, err := model.GetEvents()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"Status": "OK", "Events": events})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}

func EventView(c *gin.Context) {
	eventsTab, err := service.GetEventsTabView()
	
	if err == nil {
		c.JSON(http.StatusOK, eventsTab)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}

func EventCreate(c *gin.Context) {
	var newEvent view.NewEventFormView
	c.ShouldBindJSON(&newEvent)
	err := service.CreateEvent(&newEvent)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func EventRefresh(c *gin.Context) {
	events, err := service.LoadAndScheduleEvents()
	if err == nil {
		schedule := make(map[uint]time.Time)
		for _, event := range events {
			schedule[event.ID] = event.GetTime()
		}
		c.JSON(http.StatusOK, gin.H{"Status": "OK", "Schedule": schedule})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}
