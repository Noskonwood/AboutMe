package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func main() {
	cfg := LoadConfig()
	telegramBotToken := cfg.TelegramBotToken

	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch strings.ToLower(update.Message.Text) {
		case "/start", "/help":
			msg.Text = "Hi, I'm bot created by Bogdan Petrukhin\n Available commands:\n/about - provides short info about me\n/links - provides a list of my social links (GitHub, LinkedIn, etc)"
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("/about"),
					tgbotapi.NewKeyboardButton("/links"),
				),
			)
		case "/about":
			msg.Text = "Hi, my name is Bogdan and I'm going to be a software engineer in a future. I live in Canada and warm welcome to my telegram bot"
		case "/links":
			msg.Text = "You can find me on the following platforms:\n\nGitHub: https://github.com/Noskonwood\nLinkedIn: https://www.linkedin.com/in/bogdan-petrukhin/"
		default:
			continue
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
