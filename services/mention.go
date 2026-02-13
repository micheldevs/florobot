package services

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	i18n "github.com/micheldevs/florobot/clients/i18n"
	tg "github.com/micheldevs/florobot/clients/tg"
	"github.com/micheldevs/florobot/utils"
)

var mentions = utils.GetItemsFromCsv(fmt.Sprintf("assets/csv/%s/mentions.csv", i18n.GetCurrLang()))

func ReplyToLikelyMention(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	for _, mention := range mentions {
		match, _ := regexp.MatchString(mention[0], strings.ToLower(update.Message.Text))
		if match {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			mentionReplies := strings.Split(mention[1], ";;")
			tg.SendTxtMsg(bot, update.Message.Chat.ID, mentionReplies[r.Intn(len(mentionReplies))])
			break
		}
	}
}
