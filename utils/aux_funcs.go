package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func MatchTextInWordList(text string, wordlist []string) bool {
	for i := 0; i < len(wordlist); i++ {
		match, _ := regexp.MatchString(`(.+\s+)?`+wordlist[i]+`(\s+.+)?`, strings.ToLower(text))
		if match {
			return true
		}
	}
	return false
}

func GetCallBackKeyValue(str string) (string, string) {
	keyValue := strings.Split(str, "=")
	return keyValue[0], keyValue[1]
}

func Config(key string) string {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalln("Error loading .env file")
		}
	}

	return os.Getenv(key)
}

func GetItemsFromSingleColCsv(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '|'

	items := []string{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Error:", err)
		}

		items = append(items, record[0])
	}

	return items
}

func GetItemsFromCsv(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '|'

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Error:", err)
	}

	return records
}

func GetKeywordsNotBlacklisted(filename string, chatId int64) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '|'

	keywords := []string{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Error:", err)
		}

		if strings.Contains(record[1], strconv.FormatInt(chatId, 10)) { // the second field is the chatids separated by commas (blacklist of keywords)
			continue
		}

		keywords = append(keywords, record[0])
	}

	return keywords
}

func PrintHeader() {
	headerBytes, err := os.ReadFile("assets/app_header.txt")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	fmt.Print("\033[1;32m" + string(headerBytes) + "\033[0m\n")
}

func GetNextBgExecutionTime(lastBgTasksExecutionTime time.Time) time.Time {
	parser := cron.NewParser(
		cron.Minute |
			cron.Hour |
			cron.Dom |
			cron.Month |
			cron.Dow |
			cron.Descriptor, // allows @daily, @weekly, etc.
	)

	schedule, err := parser.Parse(Config("TG_BOT_BACKGROUND_TASKS_CRON"))
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	nextExecutionTime := schedule.Next(lastBgTasksExecutionTime)
	return nextExecutionTime

}

func CleanUpExecution() {
	log.Println("Clean up executed!")
}
