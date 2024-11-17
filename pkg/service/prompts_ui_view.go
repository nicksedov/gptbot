package service

import (
	"strconv"

	"gptbot/pkg/model"
	"gptbot/pkg/view"
)

func GetPromptsTabView(filter string) (*view.PromptsTabView, error) {

	prompts, dbErr := model.GetAll[model.Prompt]()
	if dbErr != nil {
		return nil, dbErr
	}
	promptId, filterErr := strconv.Atoi(filter)
	showFiltered := (filterErr == nil)
	promptsMap := make(map[uint]*view.PromptView, len(*prompts))
	for _, prompt := range *prompts {
		hidden := showFiltered && (prompt.ID != uint(promptId))
		promptView := view.PromptView{ID: prompt.ID, Title: prompt.Title, Prompt: prompt.Prompt, AltText: prompt.AltText, Hidden: hidden}
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
	promptViews := make([]view.PromptView, 0, len(promptsMap))
	for _, value := range promptsMap {
		promptViews = append(promptViews, *value)
	}
	return &view.PromptsTabView{PromptViews: promptViews}, nil
}
