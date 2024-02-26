package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/nicksedov/gptbot/pkg/controller"
	"github.com/nicksedov/gptbot/pkg/service"
	"github.com/nicksedov/gptbot/pkg/settings"
	"github.com/nicksedov/gptbot/pkg/telegram"
)

func main() {
	flag.Parse()
	// Init telegram bot
	_, tgErr := telegram.GetBot()
	if tgErr != nil {
		panic(tgErr)
	}
	// Init database and read events
	_, dbErr := service.ScheduleEvents()
	if dbErr != nil {
		panic(dbErr)
	}
	// Init HTTP server
	settings := settings.GetSettings()
	router := gin.Default()
	router.GET("/events/view", controller.EventView)
	router.POST("/events/create", controller.EventCreate)
	router.PUT("/events/update/:id", controller.EventUpdate)
	router.DELETE("/events/delete/:id", controller.EventDelete)

	router.GET("/prompts/view", controller.PromptView)
	router.POST("/prompts/create", controller.PromptCreate)
	router.DELETE("/prompts/delete/:id", controller.PromptDelete)

	router.GET("/chats/view", controller.ChatView)
	router.POST("/chats/create", controller.ChatCreate)
	router.DELETE("/chats/delete/:id", controller.ChatDelete)

	serverAddress := fmt.Sprintf("%s:%d", settings.Server.Host, settings.Server.Port)
	srvErr := router.Run(serverAddress)
	if srvErr != nil {
		panic(srvErr)
	}
}
