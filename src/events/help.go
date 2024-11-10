package events

import (
	tgg "api-test/tgapi"
)

func HelpHandler(message *tgg.CallbackQuery) {
	err := message.Message.SendMessage("Help request received.")
	if err != nil {
		panic(err)
	}
}
