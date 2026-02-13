package clients

import (
	"log"

	models "github.com/micheldevs/florobot/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("apps/sqlite/florobot.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error:", err)
	}

	DB.AutoMigrate(&models.Movie{}, &models.Chat{})

	return DB
}
