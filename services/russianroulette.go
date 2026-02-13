package services

import (
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	i18n "github.com/micheldevs/florobot/clients/i18n"
	tg "github.com/micheldevs/florobot/clients/tg"
)

type RussianRouletteMatch struct {
	Players        []RussianRoulettePlayer
	NextPlayerTurn int
	LastTurnTime   time.Time
}

type RussianRoulettePlayer struct {
	PlayerId       int64
	PlayerUserName string
}

var currentMatches = make(map[int64]RussianRouletteMatch) // key as chatId

func StartRussianRouletteMatch(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if _, ok := currentMatches[update.Message.Chat.ID]; ok {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulAlreadyStartedMatch1"))
		return
	}

	userName := update.Message.From.FirstName
	if update.Message.From.UserName != "" {
		userName = "@" + update.Message.From.UserName
	}

	currentMatches[update.Message.Chat.ID] = RussianRouletteMatch{[]RussianRoulettePlayer{
		{update.Message.From.ID, userName},
	}, -1, time.Now().Local()}

	tg.SendMsg(bot, tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileURL("https://64.media.tumblr.com/05c29c35b98ad86c9c4065c32340c71a/tumblr_oko5k5jRrt1vtzh7ho1_400.gif")))
	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.TransWithValues("rusRoulStartMatch1", map[string]string{"userName": userName}))
	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulStartMatch2"))
	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulStartMatch3"))
}

func JoinBotRouletteMatch(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	rsrulmatch, ok := currentMatches[update.Message.Chat.ID]
	if !ok {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulNoCurrentMatch"))
		return
	}

	if rsrulmatch.NextPlayerTurn != -1 {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulAlreadyStartedMatch3"))
		return
	}

	for _, player := range rsrulmatch.Players {
		if player.PlayerId == bot.Self.ID {
			tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulAlreadyJoinedMatch2"))
			return
		}
	}

	rsrulmatch.Players = append(rsrulmatch.Players, RussianRoulettePlayer{bot.Self.ID, "@" + bot.Self.UserName})
	rsrulmatch.LastTurnTime = time.Now().Local()
	currentMatches[update.Message.Chat.ID] = rsrulmatch

	tg.SendMsg(bot, tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileURL("https://64.media.tumblr.com/fb1ce5abf69ec6d7e0f902238accc2e3/d9f40a8ac16dfb16-af/s500x750/0a3b141a097af460be6ee5f870bedc07ce2af210.gif")))

	replies := []string{i18n.Trans("rusRoulBotJoinedMatch1"), i18n.Trans("rusRoulBotJoinedMatch2")}
	tg.SendTxtMsg(bot, update.Message.Chat.ID, replies[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(replies))])
}

func JoinRussianRouletteMatch(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	rsrulmatch, ok := currentMatches[update.Message.Chat.ID]
	if !ok {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulNoCurrentMatch"))
		return
	}

	if rsrulmatch.NextPlayerTurn != -1 {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulAlreadyStartedMatch2"))
		return
	}

	for _, player := range rsrulmatch.Players {
		if player.PlayerId == update.Message.From.ID {
			tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulAlreadyJoinedMatch1"))
			return
		}
	}

	userName := update.Message.From.FirstName
	if update.Message.From.UserName != "" {
		userName = "@" + update.Message.From.UserName
	}

	rsrulmatch.Players = append(rsrulmatch.Players, RussianRoulettePlayer{update.Message.From.ID, userName})
	rsrulmatch.LastTurnTime = time.Now().Local()
	currentMatches[update.Message.Chat.ID] = rsrulmatch

	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.TransWithValues("rusRoulJoinedMatch", map[string]string{"userName": userName}))
}

func CloseRussianRouletteMatch(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	rsrulmatch, ok := currentMatches[update.Message.Chat.ID]
	if !ok {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulNoCurrentMatch"))
		return
	}

	if rsrulmatch.NextPlayerTurn != -1 {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulCantCloseAlreadyStartedMatch"))
		return
	}

	if len(rsrulmatch.Players) < 2 {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulCantCloseMatchWithLess2Players"))
		return
	}

	rsrulmatch.NextPlayerTurn = 0
	rsrulmatch.LastTurnTime = time.Now().Local()
	currentMatches[update.Message.Chat.ID] = rsrulmatch

	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulClosePlayersMatch"))
	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulRules"))
	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.TransWithValues("rusRoulPlayerTurn", map[string]string{"userName": currentMatches[update.Message.Chat.ID].Players[0].PlayerUserName}))
}

func RollRussianRouletteMatch(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	rsrulmatch, ok := currentMatches[update.Message.Chat.ID]
	if !ok {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulNoCurrentMatch"))
		return
	}

	if rsrulmatch.NextPlayerTurn == -1 {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulDoesntClosePlayers"))
		return
	}

	if rsrulmatch.Players[rsrulmatch.NextPlayerTurn].PlayerId != update.Message.From.ID {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulNotPlayerTurn"))
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	revolverBullet := r.Intn(7)
	revolverBulletHammer := r.Intn(7)

	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulOnom1")) // https://media1.tenor.com/m/p38XIgRTGMkAAAAC/gun.gif
	tg.SendMsg(bot, tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileURL("https://i.gifer.com/JHu2.gif")))

	rsrulmatch.LastTurnTime = time.Now().Local()

	userName := update.Message.From.FirstName
	if update.Message.From.UserName != "" {
		userName = "@" + update.Message.From.UserName
	}

	if revolverBullet == revolverBulletHammer {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulOnom2"))
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.TransWithValues("rusRoulPlayerDies", map[string]string{"userName": userName}))

		rsrulmatch.Players = append(rsrulmatch.Players[:rsrulmatch.NextPlayerTurn], rsrulmatch.Players[rsrulmatch.NextPlayerTurn+1:]...)
		rsrulmatch.NextPlayerTurn = rsrulmatch.NextPlayerTurn - 1

		if len(rsrulmatch.Players) == 1 {
			if rsrulmatch.Players[0].PlayerId == bot.Self.ID {
				tg.SendMsg(bot, tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileURL("https://media1.tenor.com/m/n81BVN-hz4IAAAAd/metal-gear-solid-mgs.gif")))
				tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulBotTaunt"))
			} else {
				tg.SendMsg(bot, tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileURL("https://media1.tenor.com/m/AOpoZxBveAYAAAAd/silent-hill2-james-honey.gif")))
				tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.TransWithValues("rusRoulPlayerWins1", map[string]string{"userName": rsrulmatch.Players[0].PlayerUserName}))
			}

			delete(currentMatches, update.Message.Chat.ID)

			tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulEndMatch"))
			return
		}
	} else {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulOnom3"))
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.TransWithValues("rusRoulPlayerLives", map[string]string{"userName": userName}))
	}

	rsrulmatch.NextPlayerTurn = rsrulmatch.NextPlayerTurn + 1
	if rsrulmatch.NextPlayerTurn >= len(rsrulmatch.Players) {
		rsrulmatch.NextPlayerTurn = 0
	}

	if rsrulmatch.Players[rsrulmatch.NextPlayerTurn].PlayerId == bot.Self.ID {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.TransWithValues("rusRoulBotTurn1", map[string]string{"userName": currentMatches[update.Message.Chat.ID].Players[rsrulmatch.NextPlayerTurn].PlayerUserName}))
		time.Sleep(time.Duration(3) * time.Second)
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulBotTurn2"))

		time.Sleep(time.Duration(2) * time.Second)

		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulBotTurn3"))
		time.Sleep(time.Duration(2) * time.Second)

		revolverBullet = r.Intn(7)
		revolverBulletHammer = r.Intn(7)

		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulOnom1")) // https://media1.tenor.com/m/p38XIgRTGMkAAAAC/gun.gif
		tg.SendMsg(bot, tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileURL("https://i.gifer.com/JHu2.gif")))

		rsrulmatch.LastTurnTime = time.Now().Local()

		time.Sleep(time.Duration(2) * time.Second)
		if revolverBullet == revolverBulletHammer {
			tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulOnom2"))
			time.Sleep(time.Duration(2) * time.Second)
			replies := []string{i18n.Trans("rusRoulBotDies1"), i18n.Trans("rusRoulBotDies2")}
			tg.SendTxtMsg(bot, update.Message.Chat.ID, replies[r.Intn(len(replies))])

			rsrulmatch.Players = append(rsrulmatch.Players[:rsrulmatch.NextPlayerTurn], rsrulmatch.Players[rsrulmatch.NextPlayerTurn+1:]...)
			rsrulmatch.NextPlayerTurn = rsrulmatch.NextPlayerTurn - 1

			if len(rsrulmatch.Players) == 1 {
				tg.SendMsg(bot, tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FileURL("https://media1.tenor.com/m/AOpoZxBveAYAAAAd/silent-hill2-james-honey.gif")))
				tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.TransWithValues("rusRoulPlayerWins2", map[string]string{"userName": rsrulmatch.Players[0].PlayerUserName}))

				delete(currentMatches, update.Message.Chat.ID)

				tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulEndMatch"))
				return
			}
		} else {
			tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulOnom3"))
			tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("rusRoulBotLives"))
		}

		rsrulmatch.NextPlayerTurn = rsrulmatch.NextPlayerTurn + 1
		if rsrulmatch.NextPlayerTurn >= len(rsrulmatch.Players) {
			rsrulmatch.NextPlayerTurn = 0
		}
	}

	currentMatches[update.Message.Chat.ID] = rsrulmatch
	tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.TransWithValues("rusRoulPlayerTurn", map[string]string{"userName": currentMatches[update.Message.Chat.ID].Players[rsrulmatch.NextPlayerTurn].PlayerUserName}))
}

func CloseExpiredRollRussianRouletteMatches(bot *tgbotapi.BotAPI) {
	for chatId, rsrulmatch := range currentMatches {
		if rsrulmatch.NextPlayerTurn == -1 {
			if time.Now().After(rsrulmatch.LastTurnTime.Add(time.Duration(5) * time.Minute)) {

				delete(currentMatches, chatId)

				tg.SendTxtMsg(bot, chatId, i18n.Trans("rusRoulExpiredMatch1"))
			} else if time.Now().After(rsrulmatch.LastTurnTime.Add(time.Duration(3) * time.Minute)) {
				tg.SendTxtMsg(bot, chatId, i18n.Trans("rusRoulExpiringMatch1"))
			}
		} else {
			if time.Now().After(rsrulmatch.LastTurnTime.Add(time.Duration(10) * time.Minute)) {

				delete(currentMatches, chatId)

				tg.SendMsg(bot, tgbotapi.NewAnimation(chatId, tgbotapi.FileURL("https://i.pinimg.com/originals/85/01/18/8501189152473f6bd7a6767d84159bd1.gif")))
				tg.SendTxtMsg(bot, chatId, i18n.Trans("rusRoulExpiredMatch2"))
			} else if time.Now().After(rsrulmatch.LastTurnTime.Add(time.Duration(5) * time.Minute)) {
				tg.SendTxtMsg(bot, chatId, i18n.Trans("rusRoulExpiringMatch2"))
			}
		}
	}
}
