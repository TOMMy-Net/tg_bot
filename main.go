package main

import (
	"log"
	"os"
	"tg_bot/internal"
	"tg_bot/internal/db"
	"tg_bot/internal/tools"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	tools.NewValidator() // singletone validator
	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Bot debug for print information
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	settings := internal.NewSettings()
	settings.Bot = bot
	settings.Storage = db

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		up := update
		go proccesUpdate(&up, settings)
	}
}

func proccesUpdate(update *tgbotapi.Update, settings *internal.Settings) {

	if update.Message != nil && update.Message.Chat.Type == "private" { // If we got a message
		if settings.CheckApplication(update.Message.From.ID) {
			settings.ApplicationEnter(update)
		} else if update.Message.IsCommand() {
			settings.Commands(update)
		} else {
			settings.Messages(update)
		}
	} else if update.CallbackQuery != nil {
		settings.CallbackQuery(update)
	}
}
