package services

import (
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	i18n "github.com/micheldevs/florobot/clients/i18n"
	tg "github.com/micheldevs/florobot/clients/tg"
	"github.com/micheldevs/florobot/dao"
)

func NotificateLastMoviesToSubCineChats(bot *tgbotapi.BotAPI) {

	chats, _ := dao.GetListenableChats(-1)

	for _, chat := range chats {
		if !chat.IsSubCine {
			continue
		}

		if time.Now().AddDate(0, 0, -7).After(chat.LastNotificationCine) {
			NotificateLastMovies(bot, chat.ChatId)

			dao.UpdateNowLastNotificationCineChat(&chat)
		}
	}
}

func NotificateLastMovies(bot *tgbotapi.BotAPI, chatId int64) {
	movies, _ := dao.GetMoviesFrom(time.Date(time.Now().Local().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local),
		time.Date(time.Now().Local().Year(), time.Now().Month()+1, 1, 0, 0, 0, 0, time.Local))

	if len(movies) > 0 {
		announcement := i18n.Trans("moviesEntrance")

		nextPremiereMovies := i18n.Trans("moviesNextPremieres") + "\n\n"
		nextPremiereMoviesCounter := 0
		lastWeekPremiereMovies := i18n.Trans("moviesOldPremieres") + "\n\n"
		lastWeekPremiereMoviesCounter := 0
		pastPremiereMovies := i18n.Trans("moviesOlderPremieres") + "\n\n"
		pastPremiereMoviesCounter := 0
		for _, movie := range movies {
			movieStr := "'" + movie.Title + "' - " + movie.Duration +
				" - " + movie.PremiereDate.Format("02/01/2006") + " - " + movie.Genres +
				"\n/movdet_" + movie.ExtId + "\n"

			if movie.PremiereDate.After(time.Now()) {
				nextPremiereMoviesCounter++
				nextPremiereMovies += strconv.Itoa(nextPremiereMoviesCounter) + ". " + movieStr + "\n"
			} else if movie.PremiereDate.After(time.Now().AddDate(0, 0, -7)) {
				lastWeekPremiereMoviesCounter++
				lastWeekPremiereMovies += strconv.Itoa(lastWeekPremiereMoviesCounter) + ". " + movieStr + "\n"
			} else {
				pastPremiereMoviesCounter++
				pastPremiereMovies += strconv.Itoa(pastPremiereMoviesCounter) + ". " + movieStr + "\n"
			}
		}

		tg.SendTxtMsg(bot, chatId, announcement)
		tg.SendTxtMsg(bot, chatId, nextPremiereMovies)
		tg.SendTxtMsg(bot, chatId, lastWeekPremiereMovies)
		tg.SendTxtMsg(bot, chatId, pastPremiereMovies)
		tg.SendTxtMsg(bot, chatId, i18n.Trans("moviesConsultAgain"))
	}
}

func GetMovieDetails(bot *tgbotapi.BotAPI, update tgbotapi.Update, movieExtId string) {
	movie, err := dao.GetMovieByExtId(movieExtId)
	if err != nil {
		tg.SendTxtMsg(bot, update.Message.Chat.ID, i18n.Trans("moviesNotFound"))
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, i18n.TransWithValues("movieInfo", map[string]string{"movieName": movie.Title, "movieDuration": movie.Duration,
		"moviePremiereDate": movie.PremiereDate.Format("02/01/2006"), "movieGenres": movie.Genres, "movieDirector": movie.Director, "movieActors": movie.Actors, "movieSynopsis": movie.Synopsis, "movieExtUrl": movie.ExtUrl}))
	// clients.SendMsg(bot, tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(movie.PosterUrl)))
	tg.SendMsg(bot, msg)
}
