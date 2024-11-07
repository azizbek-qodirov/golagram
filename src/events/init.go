package events

import tgg "api-test/tgapi"

func InitializeEvents(events *tgg.Events) {

	events.Add(StartHandler, func(event *tgg.Event) bool {
		return event.Message.Text == "/start"
	})

	events.Add(HelpHandler, func(e *tgg.Event) bool {
		return e.Message.Text == "/help"
	})

	events.Add(HandleOtherMessages, func(e *tgg.Event) bool {
		return true
	})

}
