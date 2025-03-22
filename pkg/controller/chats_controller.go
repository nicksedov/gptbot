package controller

import (
	"strconv"

	"gptbot/pkg/model"
	"gptbot/pkg/service"
	"gptbot/pkg/view"

	"github.com/gin-gonic/gin"
)

func ChatView(c *gin.Context) (interface{}, error) {
	return service.GetChatsTabView()
}

func ChatCreate(c *gin.Context) (interface{}, error) {
	chatView := &view.ChatView{}
	if err := c.ShouldBindJSON(chatView); err != nil {
		return nil, err
	}
	if err := model.CreateChat(&model.TelegramChat{
		ChatID:   chatView.ChatID,
		ChatName: chatView.ChatName,
	}); err != nil {
		return nil, err
	}
	return chatView, nil
}

func ChatDelete(c *gin.Context) (interface{}, error) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 0, 0)
	if err != nil {
		return nil, err
	}
	if err := model.DeleteChat(uint(id)); err != nil {
		return nil, err
	}
	return gin.H{"status": "deleted"}, nil
}
