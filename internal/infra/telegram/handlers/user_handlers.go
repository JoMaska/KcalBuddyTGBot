package handlers

import (
	"context"
	"fmt"

	"github.com/JoMaska/KcalBuddyTGBot/internal/users/entities"
	"github.com/JoMaska/KcalBuddyTGBot/internal/users/usecases"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handlers) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Я понимаю только команды",
	})

}

func (h *Handlers) StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	input := usecases.UserRegisterInput{
		TelegramID: entities.TelegramID(update.Message.From.ID),
		Username:   update.Message.From.Username,
	}

	output, err := h.registerUser.Register(input)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Ошибка регистрации: " + err.Error(),
		})
		return
	}

	var text string
	if output.IsNew {
		text = fmt.Sprintf("Добро пожаловать, %s!\nВаш профиль создан", output.User.Username)
	} else {
		text = fmt.Sprintf("С возвращением, %s!", output.User.Username)
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   text,
	})
}
