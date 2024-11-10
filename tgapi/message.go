package tgapi

import (
	"api-test/tgapi/internal/api"
)

type Message struct {
	api       *api.Client
	token     string
	MessageID int64    `json:"message_id"`
	From      *User    `json:"from"`
	Chat      *Chat    `json:"chat"`
	Date      int64    `json:"date"`
	Text      string   `json:"text"`
	Entities  []Entity `json:"entities"`
}

func (e *Message) SendMessage(text string) error {
	return e.api.SendMessage(e.Chat.ID, text, nil)
}

func (e *Message) Reply(text string) error {
	return e.api.SendMessage(e.Chat.ID, text, &e.MessageID)
}
