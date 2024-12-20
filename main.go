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
	router.GET("/events/view", controller.EventView)
	router.POST("/events/create", controller.EventCreate)
	router.PUT("/events/update/:id", controller.EventUpdate)
	router.DELETE("/events/:id", controller.EventDelete)
	router.DELETE("/events/expired", controller.EventDeleteExpired)

	router.GET("/prompts/view", controller.PromptView)
	router.POST("/prompts/create", controller.PromptCreate)
	router.DELETE("/prompts/delete/:id", controller.PromptDelete)

	router.GET("/chats/view", controller.ChatView)
	router.POST("/chats/create", controller.ChatCreate)
	router.DELETE("/chats/delete/:id", controller.ChatDelete)

	router.POST("/messages/create", controller.MessageCreate)

	router.GET("/prebuilt/view", controller.PrebuiltMessageView)
	router.PUT("/prebuilt/update/:id", controller.PrebuiltMessageUpdate)

	settings := settings.GetSettings()
	serverAddress := fmt.Sprintf("%s:%d", settings.Server.Host, settings.Server.Port)
	srvErr := router.Run(serverAddress)
	if srvErr != nil {
		log.Fatalf("Error initializing HTTP server:\n  %s", srvErr.Error())
	}
}
