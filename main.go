package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"gptbot/pkg/controller"
	"gptbot/pkg/logger"
	"gptbot/pkg/service"
	"gptbot/pkg/settings"
	"gptbot/pkg/telegram"
)

func main() {
	// Set up config file path from command line parameter (or set default path)
	flag.Parse()
	logger.InitLog()

	// Init telegram bot
	_, tgErr := telegram.GetBot()
	if tgErr != nil {
		log.Fatalf("Error initializing telegram bot:\n  %s", tgErr.Error())
	}
	// Init database and read events
	_, dbErr := service.ScheduleEvents()
	if dbErr != nil {
		log.Fatalf("Error initializing database connection:\n  %s", dbErr.Error())
	}
	// Init HTTP server
	router := gin.Default()
	router.GET("/events/view", controller.Wrap(controller.EventView))
	router.POST("/events/create", controller.Wrap(controller.EventCreate))
	router.PUT("/events/update/:id", controller.Wrap(controller.EventUpdate))
	router.DELETE("/events/:id", controller.Wrap(controller.EventDelete))
	router.DELETE("/events/expired", controller.Wrap(controller.EventDeleteExpired))

	router.GET("/prompts/view", controller.Wrap(controller.PromptView))
	router.POST("/prompts/create", controller.Wrap(controller.PromptCreate))
	router.DELETE("/prompts/delete/:id", controller.Wrap(controller.PromptDelete))

	router.GET("/chats/view", controller.Wrap(controller.ChatView))
	router.POST("/chats/create", controller.Wrap(controller.ChatCreate))
	router.DELETE("/chats/delete/:id", controller.Wrap(controller.ChatDelete))

	router.POST("/messages/create", controller.Wrap(controller.MessageCreate))

	router.GET("/prebuilt/view", controller.Wrap(controller.PrebuiltMessageView))
	router.PUT("/prebuilt/update/:id", controller.Wrap(controller.PrebuiltMessageUpdate))

	settings := settings.GetSettings()
	serverAddress := fmt.Sprintf("%s:%d", settings.Server.Host, settings.Server.Port)
	srvErr := router.Run(serverAddress)
	if srvErr != nil {
		log.Fatalf("Error initializing HTTP server:\n  %s", srvErr.Error())
	}
}
