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
	_, dbErr := service.LoadAndScheduleEvents()
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
	serverAddress := fmt.Sprintf("%s:%d", settings.Server.Host, settings.Server.Port)
	srvErr := router.Run(serverAddress)
	if srvErr != nil {
		panic(dbErr)
	}
}
