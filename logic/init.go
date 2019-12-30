package logic

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func BotInit(token, publicUrl string) *tgbotapi.BotAPI {
	// инициализируем бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Bot name is %s", bot.Self.UserName)

	// инициализируем вэбхук
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(publicUrl + token))
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	return bot
}
