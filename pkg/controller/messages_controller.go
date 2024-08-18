package controller

import (
	"gptbot/pkg/model"
	"gptbot/pkg/telegram"
	"gptbot/pkg/view"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MessageCreate(c *gin.Context) {
	var instantMsg view.InstantMessageFormView
	c.ShouldBindJSON(&instantMsg)
	tgChat, err := model.GetChat(instantMsg.TelegramChatID)
	if err == nil {
		telegram.SendMarkdownText(instantMsg.MessageText, tgChat.ChatID)
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}