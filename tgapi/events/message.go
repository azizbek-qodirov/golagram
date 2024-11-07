package events

import (
	"fmt"
	"net/http"
	"net/url"
)

type Message struct {
	token     string
	MessageID int64    `json:"message_id"`
	From      *User    `json:"from"`
	Chat      *Chat    `json:"chat"`
	Date      int64    `json:"date"`
	Text      string   `json:"text"`
	Entities  []Entity `json:"entities"`
}

func (e *Message) SendMessage(text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s", e.token, e.From.ID, url.QueryEscape(text))
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status: %s", resp.Status)
	}
	return nil
}

func (e *Message) Reply(text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s", e.token, e.From.ID, url.QueryEscape(text))

	url += fmt.Sprintf("&reply_to_message_id=%d", e.MessageID)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to reply to message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to reply to message, status: %s", resp.Status)
	}

	return nil
}
