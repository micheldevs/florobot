package services

import (
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	i18n "github.com/micheldevs/florobot/clients/i18n"
	tg "github.com/micheldevs/florobot/clients/tg"
	"github.com/micheldevs/florobot/utils"
)

var replies = utils.GetItemsFromCsv(fmt.Sprintf("assets/csv/%s/questions.csv", i18n.GetCurrLang()))
var randomURLs = utils.GetItemsFromSingleColCsv(fmt.Sprintf("assets/csv/%s/questions_searchs.csv", i18n.GetCurrLang()))

func ReplyUser(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	reply := replies[r.Intn(len(replies))]
	switch reply[1] {
	case "G": // GoogleURL
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply[0])
		msg.ReplyToMessageID = update.Message.MessageID
		tg.SendMsg(bot, msg)
		PostRandomURL(bot, update)
	case "D": // Dunno
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply[0])
		msg.ReplyToMessageID = update.Message.MessageID
		tg.SendMsg(bot, msg)
	}
}

func PostRandomURL(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	question := url.QueryEscape(strings.ToLower(strings.ReplaceAll(update.Message.Text, "@"+bot.Self.UserName+" ", "")))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if update.Message.From.UserName != "" {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, "@"+update.Message.From.UserName+" "+fmt.Sprintf(randomURLs[r.Intn(len(randomURLs))], question))
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(randomURLs[r.Intn(len(randomURLs))], question))
		msg.ReplyToMessageID = update.Message.MessageID

		tg.SendMsg(bot, msg)
	}

}
