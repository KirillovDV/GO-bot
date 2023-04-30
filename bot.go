package main

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Создание экземпляра бота с использованием токена
	bot, err := tgbotapi.NewBotAPI("YOUR_TOKEN_HERE")
	if err != nil {
		log.Panic(err)
	}

	// Установка опций для бота
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Создание обработчика для команд и сообщений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Обработка входящих сообщений
	for update := range updates {
		// Проверка наличия текста в сообщении
		if update.Message == nil {
			continue
		}

		// Ответ на приветственные сообщения
		if strings.Contains(strings.ToLower(update.Message.Text), "hello") || strings.Contains(strings.ToLower(update.Message.Text), "hi") {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello there!")
			bot.Send(msg)
		}
	}
}
