package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	i18n "github.com/micheldevs/florobot/clients/i18n"
	tg "github.com/micheldevs/florobot/clients/tg"
	handlers "github.com/micheldevs/florobot/handlers"
	services "github.com/micheldevs/florobot/services"
	"github.com/micheldevs/florobot/utils"
)

func BackgroundTasks(bot *tgbotapi.BotAPI) {
	if utils.Config("TG_BOT_BACKGROUND_TASKS_CRON") == "" {
		log.Println("No background scheduled tasks configured!")
		return
	}

	log.Println("Initialized goroutine for background tasks...")

	for {
		log.Println("Background tasks execution...")

		log.Println("Manage Russian Roulette matches...")
		services.CloseExpiredRollRussianRouletteMatches(bot)

		log.Println("Manage last movies notifications for subbed chats...")
		services.NotificateLastMoviesToSubCineChats(bot)

		log.Println("Background tasks executed!")
		time.Sleep(time.Duration(1) * time.Minute)
	}
}

func BackgroundScheduledTasks(bot *tgbotapi.BotAPI) {
	log.Println("Initialized goroutine for scheduled background tasks...")
	var lastBgTasksExecutionTime time.Time = time.Now().Local()

	for {
		if nextBgTasksExecutionTime := utils.GetNextBgExecutionTime(lastBgTasksExecutionTime); time.Now().Local().After(nextBgTasksExecutionTime) {
			log.Println("Background scheduled tasks execution...")

			// log.Println("Scraping movies data...")
			// services.ScrapeMoviesWebBulkMovies()

			log.Println("Background scheduled tasks executed!")
			lastBgTasksExecutionTime = time.Now().Local()
		}

		time.Sleep(time.Duration(1) * time.Minute)
	}
}

func CatchSigTermSignal(c chan os.Signal) {
	<-c
	log.Println("Exit signal caught!")
	log.Println("Clean up before exiting...")
	utils.CleanUpExecution()
	log.Println("Exit!")
	os.Exit(1)
}

func main() {
	utils.PrintHeader()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go CatchSigTermSignal(c)

	i18n.Init()
	bot := tg.Init()

	go BackgroundScheduledTasks(bot)
	go BackgroundTasks(bot)
	handlers.Init(bot)
}
