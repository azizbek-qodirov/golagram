package user_events

import tgg "api-test/tgapi"

func (h *UserHandlers) DeleteHandler(message *tgg.Message) {
	exists, err := h.Storage.UM.CheckIfExists(int(message.From.ID))
	if err != nil {
		panic(err)
	}
	if !exists {
		err = message.Reply("You are not registered. Please send /start to register.", nil)
		if err != nil {
			panic(err)
		}
		return
	} else {
		user, err := h.Storage.UM.GetUser(int(message.From.ID))
		if err != nil {
			panic(err)
		}
		err = h.Storage.UM.DeleteUser(user.Telegram_id)
		if err != nil {
			panic(err)
		}
		err = message.Reply("You have been deleted from the database.", nil)
		if err != nil {
			panic(err)
		}
	}
}
