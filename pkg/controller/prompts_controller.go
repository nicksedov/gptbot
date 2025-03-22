package controller

import (
	"strconv"

	"gptbot/pkg/model"
	"gptbot/pkg/service"
	"gptbot/pkg/view"

	"github.com/gin-gonic/gin"
)

func PromptView(c *gin.Context) (interface{}, error) {
	filter := c.Query("filter")
	return service.GetPromptsTabView(filter)
}

func PromptCreate(c *gin.Context) (interface{}, error) {
	var newPrompt view.NewPromptFormView
	if err := c.ShouldBindJSON(&newPrompt); err != nil {
		return nil, err
	}

	prompt, params, err := service.BuildPromptFromCreateView(&newPrompt)
	if err != nil {
		return nil, err
	}

	if err := model.CreatePrompt(prompt); err != nil {
		return nil, err
	}

	for i := range params {
		params[i].PromptID = prompt.ID
	}

	if err := model.CreatePromptParams(params); err != nil {
		return nil, err
	}

	return prompt, nil
}

func PromptDelete(c *gin.Context) (interface{}, error) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 0, 0)
	if err != nil {
		return nil, err
	}

	if err := model.DeletePrompt(uint(id)); err != nil {
		return nil, err
	}

	return gin.H{"status": "deleted"}, nil
}