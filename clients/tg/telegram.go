package clients

import (
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/micheldevs/florobot/utils"
)

var sendRetries, _ = strconv.Atoi(utils.Config("TG_BOT_SEND_NUM_RETRIES"))

func Init() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(utils.Config("TG_BOT_APITOKEN"))
	if err != nil {
		log.Fatalln("Error:", err)
	}

	debug, _ := strconv.ParseBool(utils.Config("TG_BOT_DEBUG"))

	bot.Debug = debug
	return bot
}

func SendTxtMsg(bot *tgbotapi.BotAPI, chatId int64, text string) {
	msg := tgbotapi.NewMessage(chatId, text)
	if _, err := bot.Send(msg); err != nil {
		tries := 0
		for err != nil && tries < sendRetries {
			time.Sleep(time.Duration(10) * time.Second)
			_, err = bot.Send(msg)
			tries++
		}
	}
}

func SendMsg(bot *tgbotapi.BotAPI, msg tgbotapi.Chattable) {
	if _, err := bot.Send(msg); err != nil {
		tries := 0
		for err != nil && tries < sendRetries {
			time.Sleep(time.Duration(10) * time.Second)
			_, err = bot.Send(msg)
			tries++
		}
	}
}
