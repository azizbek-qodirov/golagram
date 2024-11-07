package tgapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type telegramBot struct {
	token     string
	events    *Events
	eventChan chan *Event
}

func NewTelegramBot(token string) (*telegramBot, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s/getMe", baseURL, token))
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response getMeResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if !response.Ok {
		return nil, errors.New("telegram API response not OK")
	}

	return &telegramBot{token: token, events: &Events{}, eventChan: make(chan *Event)}, nil
}

func (b *telegramBot) Run() error {
	offset := 0

	const numWorkers = 5
	for i := 0; i < numWorkers; i++ {
		go b.worker()
	}

	for {
		resp, err := http.Get(fmt.Sprintf("%s%s/getUpdates?offset=%d", baseURL, b.token, offset))
		if err != nil {
			return fmt.Errorf("failed to get updates: %w", err)
		}
		defer resp.Body.Close()

		var updateResponse struct {
			Ok     bool              `json:"ok"`
			Result []json.RawMessage `json:"result"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&updateResponse); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}

		if !updateResponse.Ok {
			return errors.New("telegram API response not OK")
		}

		for _, rawUpdate := range updateResponse.Result {
			var event Event
			if err := json.Unmarshal(rawUpdate, &event); err != nil {
				return fmt.Errorf("failed to unmarshal event: %w", err)
			}
			event.token = b.token

			b.eventChan <- &event
		}

		if len(updateResponse.Result) > 0 {
			lastUpdate := updateResponse.Result[len(updateResponse.Result)-1]
			var update map[string]interface{}
			if err := json.Unmarshal(lastUpdate, &update); err != nil {
				return fmt.Errorf("failed to unmarshal last update: %w", err)
			}
			offset = int(update["update_id"].(float64)) + 1
		}
	}
}

func (b *telegramBot) worker() {
	for event := range b.eventChan {
		if event.Message != nil {
			for _, e := range b.events.events {
				if e.condition(event) {
					e.handler(event)
					break
				}
			}
		}

		if event.CallbackQuery != nil {
			for _, e := range b.events.events {
				if e.condition(event) {
					e.handler(event)
					break
				}
			}
		}
	}
}

func (b *telegramBot) Close() {
	close(b.eventChan)
}

func (b *telegramBot) RegisterEvents(e *Events) {
	b.events = e
}

type botInfo struct {
	ID       int64  `json:"id"`
	IsBot    bool   `json:"is_bot"`
	Username string `json:"username"`
}

type getMeResponse struct {
	Ok     bool    `json:"ok"`
	Result botInfo `json:"result"`
}
