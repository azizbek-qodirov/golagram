package user_events

import (
	"api-test/src/storage"
	tgg "api-test/tgapi"
)

type UserHandlers struct {
	Storage *storage.Storage
}

func InitializeEvents(h *UserHandlers, events *tgg.Events) {
	events.AddMessageEvent(h.StartHandler, func(e *tgg.Message) bool {
		return e.Text == "/start"
	})
	events.AddMessageEvent(h.DeleteHandler, func(e *tgg.Message) bool {
		return e.Text == "/delete"
	})

	events.AddMessageEvent(h.EditTextHandler, func(e *tgg.Message) bool {
		return e.Text == "/edit"
	})

	events.AddMessageEvent(h.HelpHandler, func(e *tgg.Message) bool {
		return e.Text == "/help"
	})

	events.AddMessageEvent(h.HandleOtherMessages, func(e *tgg.Message) bool {
		return true
	})

}
