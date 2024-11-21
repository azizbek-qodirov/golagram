package tgapi

import (
	"api-test/tgapi/internal/api"
	"api-test/tgapi/models"
)

type Message struct {
	api             *api.Client
	token           string
	MessageID       int64    `json:"message_id"`
	Date            int64    `json:"date"`
	Chat            *Chat    `json:"chat"`
	From            *User    `json:"from"`
	ReplyTo         *Message `json:"reply_to_message"`
	EditDate        int64    `json:"edit_date"`
	Text            string   `json:"text"`
	Entities        []Entity `json:"entities"`
	CaptionEntities []Entity `json:"caption_entities"`
	// Audio                         *Audio                         `json:"audio"`
	// Document                      *Document                      `json:"document"`
	// Game                          *Game                          `json:"game"`
	// Photo                         []*PhotoSize                   `json:"photo"`
	// Sticker                       *Sticker                       `json:"sticker"`
	// Video                         *Video                         `json:"video"`
	// Voice                         *Voice                         `json:"voice"`
	// // VideoNote                     *VideoNote                     `json:"video_note"`
	// NewChatMembers []*User `json:"new_chat_members"`
	// LeftChatMember *User   `json:"left_chat_member"`
	// Caption        string  `json:"caption"`
	// Contact                       *Contact                       `json:"contact"`
	// Location                      *Location                      `json:"location"`
	// Venue                         *Venue                         `json:"venue"`
	// NewChatTitle string `json:"new_chat_title"`
	// NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo"`
	// DeleteChatPhoto       bool     `json:"delete_chat_photo"`
	// GroupChatCreated      bool     `json:"group_chat_created"`
	// SupergroupChatCreated bool     `json:"supergroup_chat_created"`
	// ChannelChatCreated    bool     `json:"channel_chat_created"`
	// MigrateToChatID       int64    `json:"migrate_to_chat_id"`
	// MigrateFromChatID     int64    `json:"migrate_from_chat_id"`
	// PinnedMessage         *Message `json:"pinned_message"`
	// Invoice                       *Invoice                       `json:"invoice"`
	// SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment"`
	// AuthorSignature  string `json:"author_signature"`
	// MediaGroupID     string `json:"media_group_id"`
	// ConnectedWebsite string `json:"connected_website"`
	// Animation                     *Animation                     `json:"animation"`
	// PassportData                  *PassportData                  `json:"passport_data"`
	// Poll                          *Poll                          `json:"poll"`
	// ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup"`
	// Dice                          *Dice                          `json:"dice"`
	// ViaBot *User `json:"via_bot"`
	// ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	// SenderChat *Chat `json:"sender_chat"`
	// VideoChatStarted              *VideoChatStarted              `json:"video_chat_started"`
	// VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended"`
	// VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`
	// MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	// VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled"`
	// IsAutomaticForward  bool `json:"is_automatic_forward"`
	// HasProtectedContent bool `json:"has_protected_content"`
	// WebAppData                    *WebAppData                    `json:"web_app_data"`
	// IsTopicMessage  bool  `json:"is_topic_message"`
	// MessageThreadID int64 `json:"message_thread_id"`
	// ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created"`
	// ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed"`
	// ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened"`
	// ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited"`
	// GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden"`
	// GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden"`
	// WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed"`
	// HasMediaSpoiler bool `json:"has_media_spoiler"`
	// ChatShared                    *ChatShared                    `json:"chat_shared"`
	// Story                         *Story                         `json:"story"`
	// Giveaway                      *Giveaway                      `json:"giveaway"`
	// GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed"`
	// GiveawayCreated               *GiveawayCreated               `json:"giveaway_created"`
	// GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners"`
	// UsersShared                   *UsersShared                   `json:"users_shared"`
	// LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options"`
	// ExternalReply                 *ExternalReplyInfo             `json:"external_reply"`
	// Quote                         *TextQuote                     `json:"quote"`
	// ForwardOrigin                 *MessageOrigin                 `json:"forward_origin"`
	// ReplyToStory                  *Story                         `json:"reply_to_story"`
	// BoostAddedd                   *ChatBoostAdded                `json:"boost_added"`
	// SenderBoostCount     int64  `json:"sender_boost_count"`
	// BusinessConnectionID string `json:"business_connection_id"`
	// SenderBusinessBot    *User  `json:"sender_business_bot"`
	// IsFromOffline        bool   `json:"is_from_offline"`
	// ChatBackgroundSet             *ChatBackground                `json:"chat_background_set"`
	// EffectID              string `json:"effect_id"`
	// ShowCaptionAboveMedia bool   `json:"show_caption_above_media"`
	// PaidMedia                     *PaidMediaInfo                 `json:"paid_media"`
	// RefundedPayment               *RefundedPayment               `json:"refunded_payment"`
}

func (e *Message) SendMessage(msg *models.MessageRequest) error {
	return e.api.SendMessage(msg)
}

func (e *Message) Reply(msg *models.MessageRequest) error {
	return e.api.SendMessage(msg)
}

func (e *Message) EditText(messageID int64, text string) error {
	return e.api.EditMessageText(e.Chat.ID, messageID, text)
}

func (e *Message) SendPhoto(photo string, caption string) {

}
