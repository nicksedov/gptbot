package model

import (
	"errors"
	"strings"
	"time"
)

// SingleEvent mappers
func (event SingleEvent) GetTime() time.Time {
	loc := time.FixedZone("", -event.TZOffset * 60)
	t := time.Time(event.Date).Add(time.Duration(event.TZOffset) * time.Minute)
	d := time.Duration(event.Time)
	return t.Add(d).In(loc)
}

func (event SingleEvent) GetResolvedPrompt() (string, error) {
	prompt := event.Prompt
	promptToResolve := prompt.Prompt
	if strings.TrimSpace(promptToResolve) != "" {
		paramsList := event.EventPromptParams
		for _, param := range paramsList {
			tag := param.PromptParam.Tag
			val := param.Value
			promptToResolve = strings.ReplaceAll(promptToResolve, "${"+tag+"}", val)
		}
		return promptToResolve, nil
	} else {
		return "", errors.New("event has invalid or empty prompt text")
	}
}

func (event SingleEvent) GetAltText() (string, error) {
	prompt := event.Prompt
	textToResolve := prompt.AltText
	if strings.TrimSpace(textToResolve) != "" {
		paramsList := event.EventPromptParams
		for _, param := range paramsList {
			tag := param.PromptParam.Tag
			val := param.Value
			textToResolve = strings.ReplaceAll(textToResolve, "${"+tag+"}", val)
		}
		return textToResolve, nil
	} else {
		return "", errors.New("event has invalid or empty fallback text")
	}
}
