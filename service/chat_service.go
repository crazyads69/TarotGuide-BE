package service

import (
	"context"
	"fmt"
	"golang_template/database"
	"golang_template/helper"
	"golang_template/schemas"
	"golang_template/util"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

// Defiane ChatService struct
type ChatService struct {
	db *gorm.DB
}

// Define method for ChatService
type IChatService interface {
	CreateChatInput(string, bool) (uint, error)
	CreateChatOutput(string, string, uint, bool) error
	GetGeneratedPrompt(string, string, string, string, []string) (string, error)
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
func (service *ChatService) CreateChatInput(input string, block bool) (uint, error) {
	chat := schemas.ChatInput{
		Message: input,
		Block:   block,
	}

	result := service.db.Table("chat_inputs").Create(&chat)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to create chat input")
		return 0, result.Error
	}
	// After insert complete must return chat id to create chat output
	result = service.db.Table("chat_inputs").Select("input_id").Last(&chat)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to get chat id")
		return 0, result.Error
	}
	return chat.InputID, nil
}

// Create chat output
func (service *ChatService) CreateChatOutput(output string, randomCards string, chatID uint, block bool) error {
	chat := schemas.Chat{
		Message:     output,
		RandomCards: randomCards,
		Block:       block,
		InputID:     chatID,
	}

	result := service.db.Table("chats").Create(&chat)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Failed to create chat output")
		return result.Error
	}

	return nil
}

// Get generated prompt from Gemini Pro API
func (service *ChatService) GetGeneratedPrompt(InputMessage string, userPrompt string, modelPrompt string, finalPrompt string, randomCards []string) (string, error) {
	var generatedPrompt string
	// Get config from .env
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Panic().Err(err).Msg("Failed to load config")
		return "", err
	}
	ctx := context.Background()
	// Initialize client
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GeminiKey))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create client")
		return "", err
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	model.SetTemperature(0.9)
	model.SetTopP(1)
	model.SetTopK(1)
	// model.SetMaxOutputTokens(2048)
	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
	}
	// Initialize the chat
	cs := model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text(userPrompt),
			},
			Role: "user",
		},

		// Initial prompt to the model
		&genai.Content{
			Parts: []genai.Part{
				genai.Text(modelPrompt),
			},
			Role: "model",
		},
	}

	questionPrompt := fmt.Sprintf(finalPrompt, InputMessage, strings.Join(randomCards, ", "), InputMessage)
	// Generate the prompt
	resp, err := cs.SendMessage(ctx, genai.Text(questionPrompt))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to send message")
		return "", err
	}
	// Decode JSON response and get parts of the response message
	// Add parts of the response message to generatedPrompt as a string
	for _, content := range resp.Candidates {
		for _, part := range content.Content.Parts {
			generatedPrompt += helper.ConvertPartToString(part)
		}
	}
	return generatedPrompt, nil
}
