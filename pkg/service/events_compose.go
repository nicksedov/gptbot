package service

import (
	"gptbot/pkg/gigachat"
	"gptbot/pkg/localai"
	"gptbot/pkg/model"
	"gptbot/pkg/settings"
	"log"
	"time"
)

func PreprocessEvent(event *model.SingleEvent) {
	prebuilt := &model.SingleEventPrebuiltMessage{
		EventID:     event.ID,
		Status:      model.Pending,
		RequestedAt: time.Now(),
	}
	message := ""
	err := model.CreatePrebuiltMessage(prebuilt)
	if err == nil {
		log.Printf("Run background preprocessing task for event with ID=%d\n", event.ID)
		message, _, err = localai.GetMessageByPrompt(event)
		if err == nil {
			prebuilt.Status = model.Created
			prebuilt.Message = message
			prebuilt.BuiltAt = time.Now()
		} else {
			prebuilt.Status = model.Failed
			log.Printf("Failed to build message with LocalAI, error returned: %s\n", err.Error())
		}
		
		if err != nil && settings.GetSettings().Fallback {
			log.Println("Falling back to process message with GigaChat")
			message, err = gigachat.GetMessageByPrompt(event)
			if err == nil {
				prebuilt.Status = model.Created
				prebuilt.Message = message
				prebuilt.BuiltAt = time.Now()
			} else {
				prebuilt.Status = model.Failed
				log.Printf("Failed to build message with GigaChat, error returned: %s\n", err.Error())
			}
		}
		
		err = model.UpdatePrebuiltMessage(prebuilt)
		if err != nil {
			log.Printf("Error finalizing preprocessing task for event with ID=%d:\n  %s\n", event.ID, err.Error())
		}
	} else {
		log.Printf("Error initializing preprocessing task for event with ID=%d:\n  %s\n", event.ID, err.Error())
	}
}
