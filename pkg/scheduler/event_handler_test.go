package scheduler

import (
	"fmt"
	"gptbot/pkg/cli"
	"gptbot/pkg/model"
	"gptbot/pkg/settings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	*cli.FlagConfig = "../../settings-test.yaml"
}

func TestHandler(t *testing.T) {
	event := model.GetDefaultTestEvent()
	event.Chat.ChatID = settings.GetSettings().Telegram.ServiceChat
	model.DeletePrebuiltMessages(event.ID)
	prebuiltMessage := &model.SingleEventPrebuiltMessage{
		ID: 1,
		EventID: event.ID,
		Status: model.Pending,
		RequestedAt: time.Now(),
	}
	err := model.CreatePrebuiltMessage(prebuiltMessage)
	assert.Nil(t, err)
	go HandleEvent(&event)
	
	// Emulate successful response from AI
	time.Sleep(time.Millisecond * 2000)
	prebuiltMessage.BuiltAt = time.Now()
	prebuiltMessage.Status = model.Created
	prebuiltMessage.Message = fmt.Sprintf("Это сообщение отправлено из unit-теста.\n"+
		"Проект: ```%s```\nМетод: ```%s```\nРасположение: ```%s```",
		"gptbot", "TestHandler()", "gptbot/pkg/scheduler/event_handler_test.go")
	model.UpdatePrebuiltMessage(prebuiltMessage)

	// Wait until HandleEvent completes (sleep duration must be greater than pendingInterval)
	time.Sleep(time.Second * 5)

}