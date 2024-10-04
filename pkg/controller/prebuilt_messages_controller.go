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
	eventIdVal := c.Query("eventId")
	eventId, err := strconv.ParseUint(eventIdVal, 0, 0)
	if err == nil {
		msg, err := model.ReadPrebuiltMessageByEventId(uint(eventId))
		if err == nil {
			result := &view.PrebuiltMessageView{
				ID: msg.ID,
				EventID: msg.EventID,
				Status:  statusMap[msg.Status],
				Message: msg.Message,
			}
			c.JSON(http.StatusOK, result)
			return
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
}

func PrebuiltMessageUpdate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 0, 0)
	if err == nil {
		var msgView view.PrebuiltMessageView
		c.ShouldBindJSON(&msgView)
		msg, err := model.GetById[model.SingleEventPrebuiltMessage](uint(id))
		if err == nil {
			msg.Message = msgView.Message
			err = model.UpdatePrebuiltMessage(msg)
		}
		if err == nil {
			c.Status(http.StatusOK)
			return
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
}
