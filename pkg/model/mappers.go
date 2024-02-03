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
		return "", errors.New("Prompt string is empty")
	}
}

func (event SingleEvent) GetAltText() string {
	prompt := event.Prompt
	return prompt.AltText
}
