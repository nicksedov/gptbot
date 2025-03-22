package controller

import (
	"gptbot/pkg/model"
	"gptbot/pkg/telegram"
	"gptbot/pkg/view"

	"github.com/gin-gonic/gin"
)

func MessageCreate(c *gin.Context) (interface{}, error) {
	var instantMsg view.InstantMessageFormView
	if err := c.ShouldBindJSON(&instantMsg); err != nil {
		return nil, err
	}

	tgChat, err := model.GetChat(instantMsg.TelegramChatID)
	if err != nil {
		return nil, err
	}

	if _, err := telegram.SendMarkdownText(tgChat.ChatID, instantMsg.MessageText); err != nil {
		return nil, err
	}

	return nil, nil
}