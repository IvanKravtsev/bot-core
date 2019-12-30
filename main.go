package main

import (
	"./logic"
	"net/http"
	"os"
)

var (
	//глобальные переменные для доступа к боту и вэбхуку
	port      = os.Getenv("PORT")       // получаем от среды выполения Heroku
	token     = os.Getenv("TOKEN")      // добавляем в настройки среды выполения Heroku, получаем от @bot_father
	publicUrl = os.Getenv("PUBLIC_URL") // добавляем в настройки среды выполения Heroku, получаем от Heroku
)

func main() {
	// заводим хероку, часть кода, отвечающая за взаимодействие с ним
	// и прослушивание запросов
	http.HandleFunc("/", logic.Handler)
	go http.ListenAndServe(":"+port, nil)

	//инициализируем бота см. ./logic/init.go
	var bot = logic.BotInit(token, publicUrl)

	// получаем сообщения
	updates := bot.ListenForWebhook("/" + bot.Token)

	// обрабатываем сообщения
	for update := range updates {
		if update.Message == nil {
			continue
		}
		// здесь и далее прописывается логика обработки сообщений

		// данная функция обрабатывает текстовые сообщения см. ./logic/logic.go
		logic.Replier(update, bot)
	}
}
