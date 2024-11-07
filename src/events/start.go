package events

import (
	tgg "api-test/tgapi"
)

func StartHandler(event *tgg.Event) {
	err := event.Message.SendMessage("Start request received.")
	if err != nil {
		panic(err)
	}
	err = event.Message.Reply("Hi!")
	if err != nil {
		panic(err)
	}
}
