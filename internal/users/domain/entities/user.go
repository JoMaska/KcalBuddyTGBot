package entities

import (
	"time"

	vo "github.com/JoMaska/KcalBuddyTGBot/internal/users/domain/vo"
	"github.com/google/uuid"
)

type User struct {
	ID         string
	TelegramID vo.TelegramID
	Username   string
	CreatedAt  time.Time
}

func NewUser(telegramID vo.TelegramID, username string) (*User, error) {
	return &User{
		ID:         uuid.NewString(),
		TelegramID: telegramID,
		Username:   username,
		CreatedAt:  time.Now().UTC(),
	}, nil
}
