package models

type MessageRequest struct {
	Text                  string   `json:"text"`
	ChatID                int64    `json:"chat_id"`
	ReplyToMessageID      int64    `json:"reply_to_message_id,omitempty"`
	ParseMode             string   `json:"parse_mode,omitempty"`
	Entities              []Entity `json:"entities,omitempty"`
	DisableWebPagePreview bool     `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool     `json:"disable_notification,omitempty"`
}

func NewMessageRequest(text string, chatID int64) *MessageRequest {
	return &MessageRequest{
		Text:   text,
		ChatID: chatID,
	}
}
