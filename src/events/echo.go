package events

import (
	tgg "api-test/tgapi"
	"fmt"
)

func HandleOtherMessages(event *tgg.Event) {
	err := event.Message.Reply(fmt.Sprintf("You sent: %s", event.Message.Text))
	if err != nil {
		panic(err)
	}
}
