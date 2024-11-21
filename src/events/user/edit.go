package user_events

import tgg "api-test/tgapi"

func (h *UserHandlers) EditTextHandler(message *tgg.Message) {
	err := message.EditText(message.ReplyTo.MessageID, "THIS MESSAGE IS EDITED!!!")
	if err != nil {
		err = message.SendMessage("Can't edit this message. Reply to bot's message or make sure it is not deleted.")
		if err != nil {
			panic(err)
		}
	}
}
