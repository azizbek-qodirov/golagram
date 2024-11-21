package events

import (
	user_events "api-test/src/events/user"
	"api-test/src/storage"
	tgg "api-test/tgapi"
)

type Handlers struct {
	User *user_events.UserHandlers
	// Admin *admin_events.AdminHandlers
}

func NewHandlers(stg *storage.Storage) *Handlers {
	return &Handlers{
		User: &user_events.UserHandlers{Storage: stg},
		// Admin: &admin_events.AdminHandlers{Storage: stg},
	}
}

func (h *Handlers) InitializeEvents(events *tgg.Events) {
	user_events.InitializeEvents(h.User, events)
	// admin_events.InitializeEvents(handlers.Admin, events)
}
