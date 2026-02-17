package handlers

import (
	"strings"

	i18n "github.com/micheldevs/florobot/clients/i18n"
	tg "github.com/micheldevs/florobot/clients/tg"
	services "github.com/micheldevs/florobot/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if strings.Contains(update.Message.Command(), "_") {
		fullCommand := strings.Split(update.Message.Command(), "_")

		switch fullCommand[0] {
		case "movdet":
			services.GetMovieDetails(bot, update, fullCommand[1])
		}
	} else {
		switch update.Message.Command() {
		case "start":
			tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.TransWithValues("startCmdText", map[string]string{"botUserName": bot.Self.UserName}))
		case "joke":
			services.TellJoke(bot, update)
		case "insult":
			tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("insultCmdText"))
		case "question":
			tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("questionCmdText"))
		case "rrstart":
			services.StartRussianRouletteMatch(bot, update)
		case "rrjoinbot":
			services.JoinBotRouletteMatch(bot, update)
		case "rrjoin":
			services.JoinRussianRouletteMatch(bot, update)
		case "rrclose":
			services.CloseRussianRouletteMatch(bot, update)
		case "rroll":
			services.RollRussianRouletteMatch(bot, update)
		case "listen":
			services.ListenUserOrChatGroup(bot, update)
		case "stoplisten":
			services.StopListenUserOrChatGroup(bot, update)
		case "movspremiere":
			services.NotificateLastMovies(bot, update.Message.Chat.ID)
		}
	}
}
