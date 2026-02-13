package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/micheldevs/florobot/dao"
	services "github.com/micheldevs/florobot/services"
)

func Init(bot *tgbotapi.BotAPI) {
	adminChats, _ := dao.GetAdminChats()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	// loop through each update.
	for update := range updates {
		if update.Poll != nil {
			continue
		}

		if update.CallbackQuery != nil {
			Callbacks(bot, update)
		} else if update.Message != nil {
			services.ManageNewChats(update, adminChats, bot)

			if update.Message.IsCommand() {
				Commands(bot, update)
			} else {
				Messages(bot, update)
			}

			services.ManageLikelyListens(bot, update)
			services.PostInListen(bot, update)
		}

	}
}
