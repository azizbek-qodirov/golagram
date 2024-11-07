package events

import (
	tgg "api-test/tgapi"
)

func HelpHandler(event *tgg.Event) {
	err := event.Message.SendMessage("Help request received.")
	if err != nil {
		panic(err)
	}
}
