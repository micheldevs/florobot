package dao

import (
	"errors"
	"time"

	models "github.com/micheldevs/florobot/models"
	"gorm.io/gorm"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InsertChat(update tgbotapi.Update) (models.Chat, error) {
	chat := models.Chat{
		ChatId: update.Message.Chat.ID,
	}

	if update.Message.Chat.Title != "" {
		chat.ChatName = update.Message.Chat.Title
	}

	chat.UserId = update.Message.From.ID
	chat.UserName = update.Message.From.UserName
	chat.UserFullName = update.Message.From.FirstName + " " + update.Message.From.LastName

	if result := DB.Create(&chat); result.Error != nil {
		return chat, result.Error
	}
	return chat, nil
}

func UpdateNowLastNotificationCineChat(chat *models.Chat) error {
	chat.LastNotificationCine = time.Now().Local()
	if result := DB.Save(&chat); result.Error != nil {
		return result.Error
	}
	return nil
}

func ExistChat(chatId int64) (bool, error) {
	var chat models.Chat
	if result := DB.Where("chat_id = ?", chatId).First(&chat); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, nil
	}
	return true, nil
}

func GetAdminChats() ([]models.Chat, error) {
	var chats []models.Chat
	if result := DB.Where("is_admin = ?", true).Find(&chats); result.Error != nil {
		return chats, result.Error
	}
	return chats, nil
}

func GetListenableChats(exceptChatId int64) ([]models.Chat, error) {
	var chats []models.Chat
	if result := DB.Where("chat_id <> ?", exceptChatId).Find(&chats); result.Error != nil {
		return chats, result.Error
	}
	return chats, nil
}

func IsChatIDAdmin(chatId int64) bool {
	var chat models.Chat
	if result := DB.Where("chat_id = ?", chatId).Where("is_admin = ?", true).First(&chat); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}
	}
	return true
}
