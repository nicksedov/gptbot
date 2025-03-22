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

func ProcessPromptParams(promptID uint, params []struct{Tag string; Title string}) error {
	// Получаем текущие параметры
	currentParams, err := model.GetPromptParamsByPromptID(promptID)
	if err != nil {
		return err
	}

	// Создаем карту для существующих параметров
	existingParams := make(map[string]*model.PromptParam)
	for i := range currentParams {
		existingParams[currentParams[i].Tag] = &currentParams[i]
	}

	// Обрабатываем новые параметры
	for _, p := range params {
		if p.Tag == "" {
			continue
		}

		if param, exists := existingParams[p.Tag]; exists {
			// Обновление существующего параметра
			param.Title = p.Title
			if err := model.UpdatePromptParam(param); err != nil {
				return err
			}
		} else {
			// Создание нового параметра
			newParam := model.PromptParam{
				Tag:      p.Tag,
				Title:    p.Title,
				PromptID: promptID,
			}
			if err := model.CreatePromptParam(&newParam); err != nil {
				return err
			}
		}
		delete(existingParams, p.Tag)
	}

	// Удаление параметров, отсутствующих в запросе
	for _, param := range existingParams {
		if err := model.DeletePromptParam(param.ID); err != nil {
			return err
		}
	}

	return nil
}