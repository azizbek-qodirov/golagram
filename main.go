package main

import (
	src_events "api-test/src/events"
	"api-test/src/utils"
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

	err = utils.SetBotCommands(mybot)
	if err != nil {
		panic(err)
	}

	events := tgg.NewEvents()

	src_events.InitializeEvents(events)

	mybot.RegisterEvents(events)

	fmt.Println("Bot is running...")
	if err := mybot.Run(); err != nil {
		panic(err)
	}
}

// https://api.telegram.org/bot6652950296:AAEKXUjZbtDs4Cu_XE40tPiECMOU_Zu5iTI/sendMessage?chat_id=1856563190&text=Hello
