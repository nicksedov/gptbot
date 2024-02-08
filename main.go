package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/nicksedov/gptbot/pkg/model"
	"github.com/nicksedov/gptbot/pkg/scheduler"
	"github.com/nicksedov/gptbot/pkg/settings"
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

	settings := settings.GetSettings()
	router := gin.Default()
    //example: router.GET("/albums", getAlbums)

    router.Run(fmt.Sprintf("%s:%d", settings.Server.Host, settings.Server.Port))
}
