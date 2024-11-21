package main

import (
	"api-test/src/config"
	"api-test/src/storage"
	"api-test/src/utils"
	"fmt"

	src_events "api-test/src/events"
	tgg "api-test/tgapi"
)

func main() {
	config := config.Load()
	mybot, err := tgg.NewTelegramBot(config.BOT_TOKEN)
	if err != nil {
		panic(err)
	}
	defer mybot.Close()

	storage, err := storage.NewPostgresStorage(config)
	if err != nil {
		panic(err)
	}

	err = utils.SetBotCommands(mybot)
	if err != nil {
		panic(err)
	}

	events := tgg.NewEvents()

	handlers := src_events.NewHandlers(storage)
	handlers.InitializeEvents(events)

	mybot.RegisterEvents(events)

	fmt.Println("Bot is running...")
	if err := mybot.Run(); err != nil {
		panic(err)
	}
}

// https://api.telegram.org/bot6652950296:AAEKXUjZbtDs4Cu_XE40tPiECMOU_Zu5iTI/sendMessage?chat_id=1856563190&text=Hello
