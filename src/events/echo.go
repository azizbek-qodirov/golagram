package events

import (
	tgg "api-test/tgapi"
	"fmt"
)

func HandleOtherMessages(message *tgg.Message) {
	err := message.Reply(fmt.Sprintf("You sent: %s", message.Text))
	fmt.Println(message.Text)
	if err != nil {
		panic(err)
	}
}
