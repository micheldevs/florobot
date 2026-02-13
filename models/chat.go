package models

import "time"

type Chat struct {
	ID                   uint   `gorm:"unique;primaryKey;autoIncrement"`
	ChatId               int64  `gorm:"chat_id"`
	ChatName             string `gorm:"chat_name"`
	UserId               int64  `gorm:"user_id"`
	UserName             string `gorm:"user_name"`
	UserFullName         string `gorm:"user_full_name"`
	IsAdmin              bool   `gorm:"is_admin;default:false"`
	IsSubCine            bool   `gorm:"is_sub_cine;default:false"`
	LastNotificationCine time.Time
	CreatedAt            time.Time
}
