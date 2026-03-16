package usecases

import "github.com/JoMaska/KcalBuddyTGBot/internal/users/entities"

type UserRegisterInput struct {
	TelegramID entities.TelegramID
	Username   string
}

type UserRegisterOutput struct {
	User *entities.User
}

type UserRepository interface {
	Save(user *entities.User) error
	FindUserByTelegramId(telegramID entities.TelegramID) (*entities.User, error)
}
