package services

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	i18n "github.com/micheldevs/florobot/clients/i18n"
	tg "github.com/micheldevs/florobot/clients/tg"
	"github.com/micheldevs/florobot/dao"
	"github.com/micheldevs/florobot/models"
)

var listens = make(map[int64]int64) // key as adminChatId, value as chatId

func ManageNewChats(update tgbotapi.Update, adminChats []models.Chat, bot *tgbotapi.BotAPI) {
	if existChat, err := dao.ExistChat(update.Message.Chat.ID); err != nil {
		panic(err)
	} else {
		if !existChat {
			chat, _ := dao.InsertChat(update)
			for _, chatAdmin := range adminChats {
				if chat.ChatName != "" {
					tg.SendTxtMsg(bot, chatAdmin.ChatId, i18n.TransWithValues("listenBotInGroup", map[string]string{"userFullName": chat.UserFullName, "userName": chat.UserName, "chatName": chat.ChatName}))
				} else {
					tg.SendTxtMsg(bot, chatAdmin.ChatId, i18n.TransWithValues("listenBotInChat", map[string]string{"userFullName": chat.UserFullName, "userName": chat.UserName}))
				}
			}
		}
	}
}

func ListenUserOrChatGroup(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if !dao.IsChatIDAdmin(update.Message.Chat.ID) {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("onlyAdminOption"))
		return
	}

	if _, ok := listens[update.Message.Chat.ID]; ok {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("chatAlreadyListen"))
		return
	}

	chats, _ := dao.GetListenableChats(update.Message.From.ID)
	var btns []tgbotapi.InlineKeyboardButton
	for i := 0; i < len(chats); i++ {
		chatName := chats[i].UserFullName + " @" + chats[i].UserName
		if chats[i].ChatName != "" {
			chatName = i18n.TransWithValues("listenGroup", map[string]string{"groupName": chats[i].ChatName})
		}

		btn := tgbotapi.NewInlineKeyboardButtonData(chatName, fmt.Sprintf("listen_chat=%d", chats[i].ChatId))
		btns = append(btns, btn)
	}

	var rows [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(btns); i += 2 {
		if i < len(btns) && i+1 < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i], btns[i+1])
			rows = append(rows, row)
		} else if i < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i])
			rows = append(rows, row)
		}
	}
	var keyboard = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, i18n.Trans("listenSelectChat"))
	msg.ReplyMarkup = keyboard
	tg.SendMsg(bot, msg)
}

func CreateListener(bot *tgbotapi.BotAPI, update tgbotapi.Update, chatIdStr string) {
	if !dao.IsChatIDAdmin(update.CallbackQuery.From.ID) {
		tg.SendTxtMsg(bot, update.CallbackQuery.From.ID, i18n.Trans("onlyAdminOption"))
		return
	}

	if _, ok := listens[update.CallbackQuery.From.ID]; ok {
		tg.SendTxtMsg(bot, update.CallbackQuery.From.ID, i18n.Trans("chatAlreadyListen"))
		return
	}

	chatId, _ := strconv.ParseInt(chatIdStr, 10, 64)

	listens[update.CallbackQuery.From.ID] = chatId
	tg.SendTxtMsg(bot, update.CallbackQuery.From.ID, i18n.Trans("listenChat"))
}

func StopListenUserOrChatGroup(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if !dao.IsChatIDAdmin(update.Message.Chat.ID) {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("onlyAdminOption"))
		return
	}

	if _, ok := listens[update.Message.Chat.ID]; !ok {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("listenNoChat"))
		return
	}

	delete(listens, update.Message.Chat.ID)

	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("listenStopChat"))
}

func ManageLikelyListens(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	for adminListenterChatId, listenChatId := range listens {
		if update.Message.Chat.ID == listenChatId {
			chatMessage := "" + update.Message.From.FirstName + " " + update.Message.From.LastName + " (@" + update.Message.From.UserName + "): "
			if update.Message.Chat.Title != "" {
				chatMessage = "'" + update.Message.Chat.Title + "' / " + update.Message.From.FirstName + " " + update.Message.From.LastName + " (@" + update.Message.From.UserName + "): "
			}

			chatMessage += update.Message.Text
			tg.SendTxtMsg(bot, adminListenterChatId, chatMessage)
		}
	}
}

func PostInListen(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if listenChatId, ok := listens[update.Message.Chat.ID]; ok {
		tg.SendTxtMsg(bot, listenChatId, update.Message.Text)
	}
}
