package usecases

import (
	"github.com/JoMaska/KcalBuddyTGBot/internal/users/domain/entities"
	"github.com/JoMaska/KcalBuddyTGBot/internal/users/domain/vo"
)

type UserRegisterInput struct {
	TelegramID vo.TelegramID
	Username   string
}

type UserRegisterOutput struct {
	User  *entities.User
	IsNew bool
}

type UserRepository interface {
	Save(user *entities.User) error
	FindUserByTelegramId(telegramID vo.TelegramID) (*entities.User, error)
}
