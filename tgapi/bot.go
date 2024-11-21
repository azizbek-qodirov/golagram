package tgapi

import (
	"api-test/tgapi/internal/api"
	"sync"

	// httpreq "api-test/tgapi/src/httpclient"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type TelegramBot struct {
	api       *api.Client
	token     string
	events    *Events
	eventChan chan *Event
	wg        sync.WaitGroup
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

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func NewTelegramBot(token string) (*TelegramBot, error) {
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

	return &TelegramBot{
		api:       api.NewClient(token),
		token:     token,
		events:    &Events{},
		eventChan: make(chan *Event, 100),
	}, nil
}

func (b *TelegramBot) Run() error {
	if b.eventChan == nil {
		return errors.New("event channel is closed")
	}

	const numWorkers = 5
	offset := 0

	for i := 0; i < numWorkers; i++ {
		b.wg.Add(1)
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
			event := emptyEvent()

			if err := json.Unmarshal(rawUpdate, &event); err != nil {
				return fmt.Errorf("failed to unmarshal event: %w", err)
			}

			if event.Message != nil {
				event.Message.api = b.api
				event.Message.token = b.token
			}
			if event.CallbackQuery != nil {
				event.CallbackQuery.api = b.api
				event.CallbackQuery.token = b.token
				if event.CallbackQuery.Message != nil {
					event.CallbackQuery.Message.api = b.api
					event.CallbackQuery.Message.token = b.token
				}
			}

			if event.EditedMessage != nil {
				event.EditedMessage.api = b.api
				event.EditedMessage.token = b.token
			}

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

func (b *TelegramBot) worker() {
	for event := range b.eventChan {
		if event.Message != nil {
			for _, e := range b.events.messageEvents {
				if e.condition(event.Message) {
					e.handler(event.Message)
					break
				}
			}
		}
		if event.CallbackQuery != nil {
			for _, e := range b.events.callbackQueryEvents {
				if e.condition(event.CallbackQuery) {
					e.handler(event.CallbackQuery)
					break
				}
			}
		}
	}
}

func (b *TelegramBot) Close() {
	close(b.eventChan)
}

func (b *TelegramBot) RegisterEvents(e *Events) {
	b.events = e
}

func (b *TelegramBot) SetBotCommands(commands []BotCommand) error {
	requestBody := map[string]interface{}{
		"commands": commands,
	}

	err := b.api.SetBotCommands(requestBody)
	if err != nil {
		return fmt.Errorf("failed to set bot commands: %w", err)
	}

	return nil
}
