package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        string
	CreatedAt time.Time
	Body      string
}

func NewMessage(body string) (Message, error) {
	if strings.TrimSpace(body) == "" {
		return Message{}, errors.New("body cannot be empty")
	}
	return Message{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		Body:      body,
	}, nil
}
