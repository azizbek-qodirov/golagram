package events

type User struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Language  string `json:"language_code"`
}

type Chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type Entity struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}
