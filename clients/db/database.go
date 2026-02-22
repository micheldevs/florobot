package clients

import (
	"fmt"
	"log"

	models "github.com/micheldevs/florobot/models"
	"github.com/micheldevs/florobot/utils"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", utils.Config("POSTGRES_HOST"), utils.Config("POSTGRES_USER"), utils.Config("POSTGRES_PASSWORD"), utils.Config("POSTGRES_DB"), utils.Config("POSTGRES_PORT"))
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error:", err)
	}

	DB.AutoMigrate(&models.Movie{}, &models.Chat{})

	return DB
}
