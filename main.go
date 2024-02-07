package main

import (
	"flag"
	"fmt"

	"github.com/nicksedov/gptbot/pkg/model"
	"github.com/nicksedov/gptbot/pkg/scheduler"
	"github.com/nicksedov/gptbot/pkg/telegram"
)

func main() {
	flag.Parse()
	events, dbErr := model.GetEvents()
	if dbErr != nil {
		panic(dbErr)
	}
	fmt.Printf("%v", events)
	_, tgErr := telegram.GetBot()
	if tgErr != nil {
		panic(tgErr)
	}
	var h *scheduler.GptChatEventHandler = &scheduler.GptChatEventHandler{}
	scheduler.Schedule(events, h)
}
