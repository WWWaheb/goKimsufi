package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var chatId int64
var bot *tgbotapi.BotAPI
var config telegramConfig

func StartBot(globalLogger *logrus.Logger, hwChan chan string, notifyChan chan string) {
	logger = globalLogger
	config = getConfig()

	bot = genBot(config.Token)

	go handleMessage(hwChan)

	go notify(notifyChan)

}

func handleMessage(hwChan chan string) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		//chatId = update.Message.Chat.ID

		if update.CallbackQuery != nil {

			if isInServList(update.CallbackQuery.Data) {
				hwChan <- update.CallbackQuery.Data
			}

			_, err = bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data))

			if err != nil {
				logger.Error("Could not send the answer to the callback query : " + err.Error())
			}

			sendMessage("searching for server type : "+update.CallbackQuery.Data, chatId)

		}

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		logger.Info("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message != nil {
			r := ""
			chatId = update.Message.Chat.ID
			switch update.Message.Text {
			case "servers", "Servers":
				r = "please choose the server type : "
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, r)
				msg.ReplyMarkup, err = genKeyboard()
				sendKeyboard(msg)
				break
			case "help", "Help":
				chatId = update.Message.Chat.ID
				r = "Commands : \nservers : Shows the list of possible servers"
				sendMessage(r, chatId)
				break
			}
		}
	}
}

func notify(notifyChan chan string) {
	for true {
		select {
		case n := <-notifyChan:
			m := "Availability for server found : \n" + n
			sendMessage(m, chatId)
			break
		}
	}
}

func genKeyboard() (tgbotapi.InlineKeyboardMarkup, error) {

	var k = tgbotapi.NewInlineKeyboardMarkup()
	var r []tgbotapi.InlineKeyboardButton
	var i = 0
	for _, v := range config.Hardware {
		r = append(r, tgbotapi.NewInlineKeyboardButtonData(v, v))
		i++
		if i == 4 {
			k.InlineKeyboard = append(k.InlineKeyboard, r)
			i = 0
			r = make([]tgbotapi.InlineKeyboardButton, 0)
		}
	}

	return k, nil
}

func isInServList(s string) bool {
	for _, v := range config.Hardware {
		if s == v {
			return true
		}
	}
	return false
}
