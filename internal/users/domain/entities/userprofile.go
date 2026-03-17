package entities

import (
	"time"

	"github.com/JoMaska/KcalBuddyTGBot/internal/users/domain/vo"
)

type UserProfile struct {
	UserID    string
	Name      string
	Age       int
	Weight    float64
	Height    int
	Gender    vo.Gender
	UpdatedAt time.Time
}

func NewUserProfile(userID string, name string, age int, weight float64, height int, gender vo.Gender) (*UserProfile, error) {
	return &UserProfile{
		UserID:    userID,
		Name:      name,
		Age:       age,
		Weight:    weight,
		Height:    height,
		Gender:    gender,
		UpdatedAt: time.Now().UTC(),
	}, nil
}
