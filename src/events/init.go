package events

import (
	tgg "api-test/tgapi"
	"database/sql"
)

type Handlers struct {
	DB *sql.DB
}

func InitializeEvents(events *tgg.Events) {
	events.AddMessageEvent(StartHandler, func(event *tgg.Message) bool {
		return event.Text == "/start"
	})

	events.AddCallbackQueryEvent(HelpHandler, func(e *tgg.CallbackQuery) bool {
		return e.Message.Text == "/help"
	})

	events.AddMessageEvent(HandleOtherMessages, func(e *tgg.Message) bool {
		return true
	})

}
