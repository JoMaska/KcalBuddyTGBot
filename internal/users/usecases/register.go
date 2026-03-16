package usecases

import (
	"github.com/JoMaska/KcalBuddyTGBot/internal/users/entities"
)

type RegisterUser struct {
	repo UserRepository
}

func NewRegisterUser(repo UserRepository) *RegisterUser {
	return &RegisterUser{repo: repo}
}

func (ru *RegisterUser) Register(input UserRegisterInput) (*UserRegisterOutput, error) {
	existing, _ := ru.repo.FindUserByTelegramId(input.TelegramID)
	if existing != nil {
		return &UserRegisterOutput{User: existing}, nil
	}

	user := entities.NewUser(input.TelegramID, input.Username)

	if err := ru.repo.Save(user); err != nil {
		return nil, err
	}
	return &UserRegisterOutput{User: user}, nil
}
