package controller

import (
	"net/http"
	"strconv"

	"gptbot/pkg/model"
	"gptbot/pkg/service"
	"gptbot/pkg/view"

	"github.com/gin-gonic/gin"
)

func ChatView(c *gin.Context) {
	chatsTab, err := service.GetChatsTabView()
	if err == nil {
		c.JSON(http.StatusOK, chatsTab)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}

func ChatCreate(c *gin.Context) {
	chatView := &view.ChatView{}
	err := c.ShouldBindJSON(chatView)
	if err == nil {
		model.CreateChat(&model.TelegramChat{ChatID: chatView.ChatID, ChatName: chatView.ChatName})
		c.JSON(http.StatusOK, chatView)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}

func ChatDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 0, 0)
	if err == nil {
		err = model.DeleteChat(uint(id))
	}
	if err == nil {
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}
