package services

import (
	"fmt"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	i18n "github.com/micheldevs/florobot/clients/i18n"
	tg "github.com/micheldevs/florobot/clients/tg"
	"github.com/micheldevs/florobot/utils"
)

var jokes = utils.GetItemsFromSingleColCsv(fmt.Sprintf("assets/csv/%s/jokes.csv", i18n.GetCurrLang()))
var audioApplauseUrls = utils.GetItemsFromSingleColCsv(fmt.Sprintf("assets/csv/%s/jokes_audios.csv", i18n.GetCurrLang()))

func TellJoke(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("jokeEntrance1"))
	time.Sleep(time.Second * time.Duration(1))
	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("jokeEntrance2"))
	time.Sleep(time.Second * time.Duration(1))
	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("jokeEntrance3"))
	time.Sleep(time.Second * time.Duration(1))
	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("jokeEntrance4"))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Second * time.Duration(2))
	tg.SendTxtMsg(bot, update.Message.Chat.ID, jokes[r.Intn(len(jokes))])

	randomReaction := r.Intn(4)
	switch randomReaction {
	case 2:
		time.Sleep(time.Second * time.Duration(3))
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("jokeExit1"))
	case 1:
		time.Sleep(time.Second * time.Duration(3))

		audioApplause := tgbotapi.NewAudio(update.Message.Chat.ID, tgbotapi.FileURL(audioApplauseUrls[r.Intn(len(audioApplauseUrls))]))
		tg.SendMsg(bot, audioApplause)
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("jokeExit2"))
	}
}
