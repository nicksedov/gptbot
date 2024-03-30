package service

import (
	"gptbot/pkg/model"
	"gptbot/pkg/view"
)

func BuildPromptFromCreateView(p *view.NewPromptFormView) (*model.Prompt, []model.PromptParam, error) {
	var prompt *model.Prompt = &model.Prompt{}
	var params []model.PromptParam = make([]model.PromptParam, 0, 3)

	prompt.Prompt = p.Prompt
	prompt.Title = p.Title
	prompt.AltText = p.AltText
	params = addPromptParam(params, p.ParamTag1, p.ParamTitle1)
	params = addPromptParam(params, p.ParamTag2, p.ParamTitle2)
	params = addPromptParam(params, p.ParamTag3, p.ParamTitle3)
	return prompt, params, nil
}

func addPromptParam(params []model.PromptParam, tag, title string) []model.PromptParam {
	if tag != "" {
		params = append(params, model.PromptParam{Tag: tag, Title: title})
	}
	return params
}
