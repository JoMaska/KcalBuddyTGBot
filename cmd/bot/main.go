package main

import (
	"log"
	"os"

	"github.com/JoMaska/KcalBuddyTGBot/internal/infra/telegram"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	token := os.Getenv("TELEGRAM_API_TOKEN")
	bot, err := telegram.NewTelegramBot(token)

	if err != nil {
		log.Panic(err)
	}
	bot.Start()
}
