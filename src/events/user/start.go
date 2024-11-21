package user_events

import (
	"api-test/src/models"
	tgg "api-test/tgapi"
	tgg_models "api-test/tgapi/models"
	"fmt"
	"time"
)

func (h *UserHandlers) StartHandler(message *tgg.Message) {
	exists, err := h.Storage.UM.CheckIfExists(int(message.From.ID))
	if err != nil {
		panic(err)
	}
	if !exists {
		newUser := &models.User{
			Telegram_id: int(message.From.ID),
			Fullname:    message.From.FirstName + " " + message.From.LastName,
			JoinedDate:  time.Unix(int64(message.Date), 0),
		}
		err = h.Storage.UM.CreateUser(newUser)
		if err != nil {
			panic(err)
		}
		msg := tgg_models.MessageRequest{
			ChatID: &message.Chat.ID,
			Text:   fmt.Sprintf("Hi %s! You are now registered. Welcome to our bot!", message.From.FirstName),
		}
		err = message.Reply(fmt.Sprintf("Hi %s! You are now registered. Welcome to our bot!", message.From.FirstName), nil)
		if err != nil {
			panic(err)
		}
	} else {
		err = message.Reply(fmt.Sprintf("Hi %s! You are already registered. Welcome back!", message.From.FirstName), nil)
		if err != nil {
			panic(err)
		}
	}
}
