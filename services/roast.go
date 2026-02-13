package services

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	i18n "github.com/micheldevs/florobot/clients/i18n"
	tg "github.com/micheldevs/florobot/clients/tg"
	"github.com/micheldevs/florobot/utils"
)

var roasts = utils.GetItemsFromCsv(fmt.Sprintf("assets/csv/%s/roast.csv", i18n.GetCurrLang()))
var gifURLs = utils.GetItemsFromSingleColCsv(fmt.Sprintf("assets/csv/%s/roast_gifs.csv", i18n.GetCurrLang()))

func RoastUser(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	roast := roasts[r.Intn(len(roasts))]
	switch roast[1] {
	case "N": // Normal
		SendRoastMsgInParts(bot, update, roast[0])
	case "I": // Imitate
		roastImitate := strings.ReplaceAll(update.Message.Text, "@"+bot.Self.UserName+" ", "")
		for _, c := range []string{"a", "e", "i", "o", "u"} {
			roastImitate = strings.ReplaceAll(roastImitate, c, "i")
		}
		for _, c := range []string{"A", "E", "I", "O", "U"} {
			roastImitate = strings.ReplaceAll(roastImitate, c, "I")
		}
		SendRoastMsgInParts(bot, update, "%s "+roastImitate)
		tg.SendTxtMsg(bot, update.Message.Chat.ID, roast[0])
	case "D": // Doxxed
		SendRoastMsgInParts(bot, update, roast[0])
		time.Sleep(time.Duration(1) * time.Second)
		tg.SendTxtMsg(bot, update.Message.Chat.ID, GetRandomDoxxingMsg())
	case "G": // GIF
		SendRoastMsgInParts(bot, update, roast[0])
		SendRandomGifMsg(bot, update.Message.Chat.ID)
	case "P": // Poll
		userName := update.Message.From.FirstName
		if update.Message.From.UserName != "" {
			userName = "@" + update.Message.From.UserName
		}
		poll := tgbotapi.NewPoll(update.Message.Chat.ID, fmt.Sprintf(roast[0], userName), "Sí", "No")
		tg.SendMsg(bot, poll)
	}
}

func SendRoastMsgInParts(bot *tgbotapi.BotAPI, update tgbotapi.Update, roastMsg string) {
	roasts := strings.Split(roastMsg, "\n")
	for _, roast := range roasts {
		time.Sleep(time.Duration(1) * time.Second)
		if strings.Contains(roast, "%s") {
			if update.Message.From.UserName != "" {
				tg.SendTxtMsg(bot, update.Message.Chat.ID, fmt.Sprintf(roast, "@"+update.Message.From.UserName))
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(roast, update.Message.From.FirstName))
				msg.ReplyToMessageID = update.Message.MessageID
				tg.SendMsg(bot, msg)
			}
		} else {
			tg.SendTxtMsg(bot, update.Message.Chat.ID, roast)
		}
	}
}

func GetRandomDoxxingMsg() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var doxxings = []string{
		fmt.Sprintf("http://maps.google.com/maps?z=12&t=m&q=loc:%.6f+%.6f", r.Float32()*float32(r.Intn(91)-90), r.Float32()*float32(r.Intn(181)-180)),
		i18n.TransWithValues("roastDoxxedInfo1", map[string]string{"roastIp": fmt.Sprintf("%d.%d.%d.%d", r.Intn(256), r.Intn(256), r.Intn(256), r.Intn(256)),
			"roastDns": fmt.Sprintf("%d.%d.%d.%d", r.Intn(9), r.Intn(9), r.Intn(9), r.Intn(9)), "roastW": fmt.Sprintf("%d %04d", r.Intn(100), r.Intn(10000)),
			"roastN": fmt.Sprintf("%d %04d", r.Intn(100), r.Intn(10000))}),
		i18n.TransWithValues("roastDoxxedInfo2", map[string]string{
			"roastCardNumber":  fmt.Sprintf("%04d %04d %04d %04d", r.Intn(10000), r.Intn(10000), r.Intn(10000), r.Intn(10000)),
			"roastCardExpDate": fmt.Sprintf("%02d/%02d", r.Intn(13), time.Now().Year()+r.Intn(6)),
			"roastCardCVV":     fmt.Sprintf("%03d", r.Intn(1000)),
			"roastIdCard":      fmt.Sprintf("%09dF", r.Intn(999999999-100000000)+100000000),
		}),
	}
	return doxxings[r.Intn(len(doxxings))]
}

func SendRandomGifMsg(bot *tgbotapi.BotAPI, chatID int64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	tg.SendMsg(bot, tgbotapi.NewAnimation(chatID, tgbotapi.FileURL(gifURLs[r.Intn(len(gifURLs))])))
}

func GetRoastsWordlist() []string {
	return utils.GetItemsFromSingleColCsv(fmt.Sprintf("assets/csv/%s/roast_keywords.csv", i18n.GetCurrLang()))
}
