package handlers

import (
	"github.com/JoMaska/KcalBuddyTGBot/internal/users/usecases"
	"github.com/go-telegram/bot"
)

type Handlers struct {
	registerUser *usecases.RegisterUser
}

func NewHandlers(registerUser *usecases.RegisterUser) (*Handlers, error) {
	return &Handlers{registerUser: registerUser}, nil
}

func (h *Handlers) RegisterAll(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, h.StartHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeContains, h.DefaultHandler)
}
