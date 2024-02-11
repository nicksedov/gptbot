package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nicksedov/gptbot/pkg/service"
)

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
