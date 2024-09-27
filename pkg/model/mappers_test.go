package model

import (
	"log"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetResolvedPrompt(t *testing.T) {
	event := GetDefaultTestEvent()
	resolved, err := event.GetResolvedPrompt()
	if err != nil {
		log.Fatal(err.Error())
	}
	expected := event.Prompt.Prompt
	pp := event.EventPromptParams
	expected = strings.ReplaceAll(expected, "${"+pp[0].PromptParam.Tag+"}", pp[0].Value)
	expected = strings.ReplaceAll(expected, "${"+pp[1].PromptParam.Tag+"}", pp[1].Value)
	assert.Equal(t, expected, resolved)
}

func TestGetTime(t *testing.T) {
	event := GetDefaultTestEvent()
	eventTime := event.GetTime()
	expectedLoc := time.FixedZone("", -tzOffset * 60)
	expectedTime := time.Date(testDate.Year(), testDate.Month(), testDate.Day(), hour, min, sec, ns, expectedLoc)
	assert.Equal(t, expectedTime, eventTime)
}
