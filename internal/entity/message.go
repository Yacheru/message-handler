package entity

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Message string `json:"message" binding:"required"`
}

type DBMessage struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Message   string    `json:"message" db:"message"`
	Marked    bool      `json:"marked,omitempty" db:"marked"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

type Statistic struct {
	Unmarked int `json:"unmarked"`
	Marked   int `json:"marked"`
}
