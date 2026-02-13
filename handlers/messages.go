package handlers

import (
	"fmt"
	"regexp"
	"strings"

	i18n "github.com/micheldevs/florobot/clients/i18n"
	"github.com/micheldevs/florobot/services"
	"github.com/micheldevs/florobot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	regexQuestion, _ := regexp.Compile(`[\w\s.,:;\-+@]*¿[\w\s,]*?`)

	if utils.MatchTextInWordList(update.Message.Text, utils.GetKeywordsNotBlacklisted(fmt.Sprintf("assets/csv/%s/jokes_keywords.csv", i18n.GetCurrLang()), update.Message.Chat.ID)) {
		services.TellJoke(bot, update)
	} else if strings.Contains(strings.ToLower(update.Message.Text), strings.ToLower("@"+bot.Self.UserName)) && utils.MatchTextInWordList(update.Message.Text, services.GetRoastsWordlist()) {
		services.RoastUser(bot, update)
	} else if strings.Contains(strings.ToLower(update.Message.Text), strings.ToLower("@"+bot.Self.UserName)) && regexQuestion.MatchString(update.Message.Text) {
		services.ReplyUser(bot, update)
	} else {
		services.ReplyToLikelyMention(bot, update)
	}
}
