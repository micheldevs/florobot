package dao

import (
	database "github.com/micheldevs/florobot/clients/db"

	"gorm.io/gorm"
)

var DB *gorm.DB = database.Init()
