package models

import "time"

type User struct {
	Telegram_id int       `json:"telegram_id"`
	Fullname    string    `json:"fullname"`
	JoinedDate  time.Time `json:"joined_date"`
}
