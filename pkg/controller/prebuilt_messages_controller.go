package controller

import (
	"gptbot/pkg/model"
	"gptbot/pkg/view"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var statusMap map[model.MessageBuildStatus]string = map[model.MessageBuildStatus]string{
	model.Pending: "pending",
	model.Created: "created",
	model.Failed:  "failed",
}

func PrebuiltMessageView(c *gin.Context) {
	eventIdVal := c.Query("id")
	eventId, err := strconv.Atoi(eventIdVal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}

	msg, err := model.ReadPrebuiltMessage(uint(eventId))
	if err != nil || msg == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}

	result := &view.PrebuiltMessageView{
		EventID: msg.EventID,
		Status:  statusMap[msg.Status],
		Message: msg.Message,
	}
	c.JSON(http.StatusOK, result)
}
