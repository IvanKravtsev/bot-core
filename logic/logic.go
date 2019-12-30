package logic

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
	"reflect"
)

func Replier(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

		var chatID = update.Message.Chat.ID
		var text = update.Message.Text

		switch {
		//в каждом кейсе логика обработки команды
		case "/start" == text:
			Start(chatID, bot)

		case "/hello" == text:
			ReplyHello(update.Message, bot)

		default:
			{
				MirrorPush(chatID, bot, text)
			}
		}
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Главная страница bot-core. Сюда можно вставить http/css код с информацией о боте"))
}

//стартовое сообщение, выводится при первом контакте с ботом, а так же после команды /start
func Start(id int64, bt *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(id, "Привет! Я — ядро для бота!")
	bt.Send(msg)
}

//reply-to-message сообщение, выводится после команды /hello
func ReplyHello(message *tgbotapi.Message, bt *tgbotapi.BotAPI) {
	var (
		id        = message.Chat.ID
		messageID = message.MessageID
	)

	msg := tgbotapi.NewMessage(id, "Привет!")
	msg.ReplyToMessageID = messageID
	bt.Send(msg)
}

//повторяет любое текстовое сообщение
func MirrorPush(id int64, bt *tgbotapi.BotAPI, txt string) {
	var pushString = "U said: " + txt
	bt.Send(tgbotapi.NewMessage(id, pushString))
}
