package handlers

import (
	services "github.com/micheldevs/florobot/services"
	utils "github.com/micheldevs/florobot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmd, objectId := utils.GetCallBackKeyValue(update.CallbackQuery.Data)
	switch cmd {
	case "listen_chat":
		services.CreateListener(bot, update, objectId)
	}
}
