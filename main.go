package main

import (
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	for true {
		updates := bot.GetUpdatesChan(updateConfig)
		for update := range updates {
			if update.Message == nil {
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No entiendo!")

			if strings.ToLower(update.Message.Text) == "casas" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, FindHouses())
			}
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
