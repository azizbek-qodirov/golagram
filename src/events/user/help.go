package user_events

import (
	tgg "api-test/tgapi"
)

func (h *UserHandlers) HelpHandler(message *tgg.Message) {
	err := message.SendMessage("Help request received.")
	if err != nil {
		panic(err)
	}
}
