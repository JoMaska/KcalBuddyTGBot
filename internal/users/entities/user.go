package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         string
	TelegramID TelegramID
	Username   string
	CreatedAt  time.Time
}

func NewUser(telegramID TelegramID, username string) (*User, error) {
	return &User{
		ID:         uuid.NewString(),
		TelegramID: telegramID,
		Username:   username,
		CreatedAt:  time.Now().UTC(),
	}, nil
}
