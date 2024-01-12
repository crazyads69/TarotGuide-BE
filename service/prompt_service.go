package service

import (
	"golang_template/database"
	"golang_template/schemas"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Define struct for PromptService
type PromptService struct {
	db *gorm.DB
}

// Define method for PromptService
type IPromptService interface {
	GetFinalPrompt() (string, error)
	GetUserPrompt() (string, error)
	GetModelPrompt() (string, error)
}

// Implement interface for PromptService struct
func PromptServiceImpl() IPromptService {
	db := database.DBConnect()
	err := db.AutoMigrate(&schemas.Prompt{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate database")
	}

	return &PromptService{
		db: db,
	}
}

// Get final prompt
func (service *PromptService) GetFinalPrompt() (string, error) {
	var finalPrompt schemas.Prompt
	result := service.db.Table("prompts").Where("prompt_id = ?", "final_prompt").First(&finalPrompt)

	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to get final prompt")
		return "", result.Error
	}

	return finalPrompt.Prompt, nil
}

// Get user prompt
func (service *PromptService) GetUserPrompt() (string, error) {
	var userPrompt schemas.Prompt
	result := service.db.Table("prompts").Where("prompt_id = ?", "user_prompt").First(&userPrompt)

	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to get user prompt")
		return "", result.Error
	}

	return userPrompt.Prompt, nil
}

// Get model prompt
func (service *PromptService) GetModelPrompt() (string, error) {
	var modelPrompt schemas.Prompt
	result := service.db.Table("prompts").Where("prompt_id = ?", "model_prompt").First(&modelPrompt)

	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to get model prompt")
		return "", result.Error
	}

	return modelPrompt.Prompt, nil
}
