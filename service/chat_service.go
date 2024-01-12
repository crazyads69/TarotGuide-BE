package service

import (
	"golang_template/database"
	"golang_template/schemas"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Defiane ChatService struct
type ChatService struct {
	db *gorm.DB
}

// Define method for ChatService
type IChatService interface {
	CreateChatInput(string, bool) error
	CreateChatOutput(string, string, bool) error
}

// Define interface for ChatService
func ChatServiceImpl() IChatService {
	db := database.DBConnect()
	err := db.AutoMigrate(&schemas.Chat{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate database")
	}

	return &ChatService{
		db: db,
	}
}

// Create chat input
func (service *ChatService) CreateChatInput(input string, block bool) error {
	chat := schemas.ChatInput{
		Message: input,
		Block:   block,
	}

	result := service.db.Table("chat_inputs").Create(&chat)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to create chat input")
		return result.Error
	}

	return nil
}

// Create chat output
func (service *ChatService) CreateChatOutput(output string, randomCards string, block bool) error {
	chat := schemas.Chat{
		Message:     output,
		RandomCards: randomCards,
		Block:       block,
	}

	result := service.db.Table("chats").Create(&chat)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to create chat output")
		return result.Error
	}

	return nil
}
