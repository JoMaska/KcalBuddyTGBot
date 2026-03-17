package telegram

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/JoMaska/KcalBuddyTGBot/internal/infra/telegram/handlers"
	"github.com/JoMaska/KcalBuddyTGBot/internal/users/domain/entities"
	"github.com/JoMaska/KcalBuddyTGBot/internal/users/domain/vo"
	"github.com/JoMaska/KcalBuddyTGBot/internal/users/usecases"
	"github.com/go-telegram/bot"
)

type TelegramBot struct {
	bot *bot.Bot
}

type userRepository struct {
}

// TODO: убрать эти заглушки
func (ur *userRepository) Save(user *entities.User) error {
	return nil
}

// TODO: убрать эти заглушки
func (ur *userRepository) FindUserByTelegramId(telegramID vo.TelegramID) (*entities.User, error) {
	return nil, nil
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	userRepository := &userRepository{}

	registerUser := usecases.NewRegisterUser(userRepository)
	h, _ := handlers.NewHandlers(registerUser)

	bot, err := bot.New(token)

	if err != nil {
		log.Panic(err)
	}

	h.RegisterAll(bot)

	return &TelegramBot{bot: bot}, nil
}

func (tgb *TelegramBot) Start() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	tgb.bot.Start(ctx)
}
