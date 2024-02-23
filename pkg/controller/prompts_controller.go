package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicksedov/gptbot/pkg/service"
)

func PromptView(c *gin.Context) {
	promtsTab, err := service.GetPromptsTabView()
	if err == nil {
		c.JSON(http.StatusOK, promtsTab)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}

func PromptCreate(c *gin.Context) {
}

func PromptDelete(c *gin.Context) {
}