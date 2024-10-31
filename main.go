package main

import (
	"log"
	"os"
	"tg_bot/internal"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	settings := internal.NewSettings()
	settings.Bot = bot

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		go proccesUpdate(&update, settings)
	}
}

func proccesUpdate(update *tgbotapi.Update, settings *internal.Settings) {
	if update.Message != nil && update.Message.Chat.Type == "private" { // If we got a message

		if update.Message.IsCommand() {
			settings.Commands(update)
		} else {
			settings.Messages(update)
		}

	} else if update.CallbackQuery != nil {
		settings.CallbackQuery(update)
	}
}
