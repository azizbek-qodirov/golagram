package main

import (
	src_events "api-test/src/events"
	"fmt"

	tgg "api-test/tgapi"
)

const BOT_TOKEN = "6652950296:AAEKXUjZbtDs4Cu_XE40tPiECMOU_Zu5iTI"

func main() {
	mybot, err := tgg.NewTelegramBot(BOT_TOKEN)
	if err != nil {
		panic(err)
	}

	defer mybot.Close()

	events := tgg.NewEvents()

	src_events.InitializeEvents(events)

	mybot.RegisterEvents(events)

	fmt.Println("Bot is running...")
	if err := mybot.Run(); err != nil {
		panic(err)
	}
}
