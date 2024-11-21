package user_events

import (
	tgg "api-test/tgapi"
	"fmt"
)

func (h *UserHandlers) HandleOtherMessages(message *tgg.Message) {
	err := message.Reply(fmt.Sprintf("You sent: %s", message.Text), nil)
	fmt.Println(message.Text)
	if err != nil {
		panic(err)
	}
}
