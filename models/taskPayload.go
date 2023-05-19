package models

type TaskPayload struct {
	Title    string `json:"title"`
	Todo     string `json:"todo"`
	Priority int    `json:"priority"`
	UserId   uint   `json:"userId"` // foreign key
}
