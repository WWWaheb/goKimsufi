package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func genBot(t string) *tgbotapi.BotAPI {
	var err error

	bot, err = tgbotapi.NewBotAPI(t)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	logger.Info("Authorized on account %s", bot.Self.UserName)

	return bot
}

func sendMessage(m string, id int64) {
	msg := tgbotapi.NewMessage(id, m)
	_, err := bot.Send(msg)
	if err != nil {
		logger.Error("Could not send message : " + err.Error())
	}
}

func sendKeyboard(m tgbotapi.MessageConfig) {
	_, err := bot.Send(m)
	if err != nil {
		logger.Error("Could not send keyboard : " + err.Error())
	}
}
