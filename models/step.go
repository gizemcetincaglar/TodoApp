package models

import "time"

type Step struct {
	ID        string     `json:"id"`
	TodoID    string     `json:"todo_id"`
	Content   string     `json:"content"`
	Completed bool       `json:"completed"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
