package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nicksedov/gptbot/pkg/model"
	"github.com/nicksedov/gptbot/pkg/service"
	"github.com/nicksedov/gptbot/pkg/view"
)

func PromptView(c *gin.Context) {
	var filter string = c.Query("filter")
	promtsTab, err := service.GetPromptsTabView(filter)
	if err == nil {
		c.JSON(http.StatusOK, promtsTab)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}

func PromptCreate(c *gin.Context) {
	var newPrompt view.NewPromptFormView
	c.ShouldBindJSON(&newPrompt)
	prompt, params, err := service.BuildPromptFromCreateView(&newPrompt)
	if err == nil {
		err = model.CreatePrompt(prompt)
		if err == nil {
			for i := 0; i < len(params); i++ {
				params[i].PromptID = prompt.ID
			}
			model.CreatePromptParams(params)
		}
	}
	onPromptsChanged(c, err)
}

func PromptDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 0, 0)
	if err == nil {
		err = model.DeletePrompt(uint(id))
	}
	onPromptsChanged(c, err)
}

func onPromptsChanged(c *gin.Context, err error) {
	if err == nil {
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Error", "ErrorMessage": err.Error()})
	}
}
