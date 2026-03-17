package vo

import (
	"errors"

	"github.com/google/uuid"
)

type UserID struct {
	ID string
}

func NewUserID() UserID {
	return UserID{ID: uuid.NewString()}
}

func ParseUserID(id string) (UserID, error) {
	if _, err := uuid.Parse(id); err != nil {
		return UserID{}, errors.New("Invalid user id")
	}
	return UserID{ID: id}, nil
}
