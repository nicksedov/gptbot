package controller

import (
	"gptbot/pkg/model"
	"gptbot/pkg/view"
	"strconv"

	"github.com/gin-gonic/gin"
)

var statusMap = map[model.MessageBuildStatus]string{
	model.Pending: "pending",
	model.Created: "created",
	model.Failed:  "failed",
}

func PrebuiltMessageView(c *gin.Context) (interface{}, error) {
	eventIdVal := c.Query("eventId")
	eventId, err := strconv.ParseUint(eventIdVal, 0, 0)
	if err != nil {
		return nil, err
	}

	msg, err := model.ReadPrebuiltMessageByEventId(uint(eventId))
	if err != nil {
		return nil, err
	}

	return &view.PrebuiltMessageView{
		ID:      msg.ID,
		EventID: msg.EventID,
		Status:  statusMap[msg.Status],
		Message: msg.Message,
	}, nil
}

func PrebuiltMessageUpdate(c *gin.Context) (interface{}, error) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 0, 0)
	if err != nil {
		return nil, err
	}

	var msgView view.PrebuiltMessageView
	if err := c.ShouldBindJSON(&msgView); err != nil {
		return nil, err
	}

	msg, err := model.GetById[model.SingleEventPrebuiltMessage](uint(id))
	if err != nil {
		return nil, err
	}

	msg.Message = msgView.Message
	if err := model.UpdatePrebuiltMessage(msg); err != nil {
		return nil, err
	}

	return gin.H{"status": "updated"}, nil
}