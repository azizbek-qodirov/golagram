package tgapi

import (
	"api-test/tgapi/internal/api"
	"fmt"
	"net/http"
	"net/url"
)

type CallbackQuery struct {
	api     *api.Client
	token   string
	ID      string   `json:"id"`
	From    *User    `json:"from"`
	Message *Message `json:"message"`
	Data    string   `json:"data"`
}

func (e *CallbackQuery) SendMessage(text string) error {
	url := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=%s", baseURL, e.token, e.From.ID, url.QueryEscape(text))
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

func (e *CallbackQuery) Reply(text string) error {
	url := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=%s", baseURL, e.token, e.From.ID, url.QueryEscape(text))
	url += fmt.Sprintf("&reply_to_message_id=%d", e.Message.MessageID)

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
