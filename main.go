package main

import (
	"flag"
	"fmt"
	"github.com/nicksedov/gptbot/pkg/database"
	"github.com/nicksedov/gptbot/pkg/telegram"
)

func main() {
	flag.Parse()
	events, dbErr := database.GetEvents()
	if dbErr != nil {
		panic(dbErr)
	}
	fmt.Printf("%v", events)
	_, tgErr := telegram.GetBot()
	if tgErr != nil {
		panic(tgErr)
	}
}
