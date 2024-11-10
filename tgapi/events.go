package tgapi

type Event struct {
	UpdateID      int64          `json:"update_id"`
	Message       *Message       `json:"message"`
	CallbackQuery *CallbackQuery `json:"callback_query"`
	EditedMessage *Message       `json:"edited_message"`
}

type Events struct {
	messageEvents []struct {
		handler   func(*Message)
		condition func(*Message) bool
	}
	callbackQueryEvents []struct {
		handler   func(*CallbackQuery)
		condition func(*CallbackQuery) bool
	}
}

func NewEvents() *Events {
	return &Events{}
}

func (e *Events) AddMessageEvent(event func(*Message), condition func(*Message) bool) {
	e.messageEvents = append(e.messageEvents, struct {
		handler   func(*Message)
		condition func(*Message) bool
	}{handler: event, condition: condition})
}

func (e *Events) AddCallbackQueryEvent(event func(*CallbackQuery), condition func(*CallbackQuery) bool) {
	e.callbackQueryEvents = append(e.callbackQueryEvents, struct {
		handler   func(*CallbackQuery)
		condition func(*CallbackQuery) bool
	}{handler: event, condition: condition})
}
