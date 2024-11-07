package tgapi

import (
	"api-test/tgapi/events"
)

type Event struct {
	token         string
	UpdateID      int64                 `json:"update_id"`
	Message       *events.Message       `json:"message"`
	CallbackQuery *events.CallbackQuery `json:"callback_query"`
	EditedMessage *events.Message       `json:"edited_message"`
}

type Events struct {
	events []struct {
		handler   func(*Event)
		condition func(*Event) bool
	}
}

func NewEvents() *Events {
	return &Events{}
}

func (e *Events) Add(event func(*Event), condition func(*Event) bool) {
	e.events = append(e.events, struct {
		handler   func(*Event)
		condition func(*Event) bool
	}{handler: event, condition: condition})
}
