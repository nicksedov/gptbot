package scheduler

import (
	"errors"
	"gptbot/pkg/model"
	"gptbot/pkg/telegram"
	"time"
)

const errorCountdownThreshold = 5
const pendingInterval = time.Second * 3
const pendingDurationThreshold = time.Minute * 5

func HandleEvent(event *model.SingleEvent) error {
	var msg *model.SingleEventPrebuiltMessage
	var err error
	var errCountdown int = errorCountdownThreshold
	for isPending(msg) && errCountdown > 0 {
		msg, err = model.ReadPrebuiltMessage(event.ID)
		if err != nil {
			errCountdown--
		} else {
			errCountdown = errorCountdownThreshold
		}
		time.Sleep(pendingInterval)
	}
	if msg != nil {
		return messageDispatch(event.Chat.ChatID, msg)
	} else {
		return err
	}
}

func isPending(msg *model.SingleEventPrebuiltMessage) bool {
	if msg == nil {
		return true
	}
	if msg.Status == model.Pending {
		pendingDuration := time.Since(msg.RequestedAt)
		return pendingDuration < pendingDurationThreshold
	} else {
		return false
	}
}

func messageDispatch(chatId int64, msg *model.SingleEventPrebuiltMessage) error {
	var err error
	if msg.Status == model.Created {
		msgText := msg.Message
		if msgText != "" {
			_, err = telegram.SendMarkdownText(chatId, msgText)
		} else {
			err = errors.New("completions API returned empty response")
		}
	} else if msg.Status == model.Pending {
		err = errors.New("completions API pending timeout")
	} else {
		err = errors.New("completions API failure")
	}
	return err
}
