package tgapi

const baseURL = "https://api.telegram.org/bot"

func emptyEvent() Event {
	return Event{
		Message: &Message{
			From: &User{},
			Chat: &Chat{},
		},
		CallbackQuery: &CallbackQuery{
			Message: &Message{
				From: &User{},
				Chat: &Chat{},
			},
			From: &User{},
		},
		EditedMessage: &Message{
			From: &User{},
			Chat: &Chat{},
		},
	}
}
