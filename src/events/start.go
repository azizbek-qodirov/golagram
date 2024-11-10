package events

import (
	tgg "api-test/tgapi"
)

func StartHandler(message *tgg.Message) {
	err := message.SendMessage("Start request received.")
	if err != nil {
		panic(err)
	}
	err = message.Reply("Hi!")
	if err != nil {
		panic(err)
	}
}
