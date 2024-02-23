package service

import (
	"github.com/nicksedov/gptbot/pkg/model"
	"github.com/nicksedov/gptbot/pkg/view"
)

func GetPromptsTabView() (*view.PromptsTabView, error) {

	prompts, dbErr := model.GetAll[model.Prompt]()
	if dbErr != nil {
		return nil, dbErr
	}
	promptsMap := make(map[uint]*view.PromptFormView, len(*prompts))
	for _, prompt := range *prompts {
		promptView := view.PromptFormView{Title: prompt.Title, Prompt: prompt.Prompt, AltText: prompt.AltText}
		promptsMap[prompt.ID] = &promptView
	}

	promptParams, dbErr := model.GetAll[model.PromptParam]()
	if dbErr != nil {
		return nil, dbErr
	}
	for _, promptParam := range *promptParams {
		id := promptParam.PromptID
		if promptsMap[id] != nil {
			pp := view.PromptParam{Tag: promptParam.Tag, Title: promptParam.Title}
			(*promptsMap[id]).PromptParams = append((*promptsMap[id]).PromptParams, pp)
		}
	}
	promptViews := make([]view.PromptFormView, 0, len(promptsMap))
	for  _, value := range promptsMap {
		promptViews = append(promptViews, *value)
	}
	return &view.PromptsTabView{PromptViews: promptViews}, nil
}
