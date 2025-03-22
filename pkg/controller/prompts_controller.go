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

func PromptUpdate(c *gin.Context) (interface{}, error) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 0, 0)
	if err != nil {
		return nil, err
	}

	var updPrompt view.UpdatePromptFormView
	if err := c.ShouldBindJSON(&updPrompt); err != nil {
		return nil, err
	}

	// Получаем промпт с параметрами
	prompt, err := model.GetPromptWithParams(uint(id))
	if err != nil {
		return nil, err
	}

	// Обновление основных полей промпта
	if updPrompt.Title != "" {
		prompt.Title = updPrompt.Title
	}
	if updPrompt.Prompt != "" {
		prompt.Prompt = updPrompt.Prompt
	}
	if updPrompt.AltText != "" {
		prompt.AltText = updPrompt.AltText
	}

	// Обновление основного промпта
	if err := model.UpdatePrompt(prompt); err != nil {
		return nil, err
	}

	// Обработка параметров
	params := []struct {
		Tag   string
		Title string
	}{
		{updPrompt.ParamTag1, updPrompt.ParamTitle1},
		{updPrompt.ParamTag2, updPrompt.ParamTitle2},
		{updPrompt.ParamTag3, updPrompt.ParamTitle3},
	}

	// Обновляем параметры через сервис
	if err := service.ProcessPromptParams(prompt.ID, params); err != nil {
		return nil, err
	}

	// Возвращаем обновленный промпт с параметрами
	return model.GetPromptWithParams(uint(id))
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