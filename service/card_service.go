package service

import (
	"golang_template/database"
	"golang_template/schemas"
	"math/rand"
	"time"

	"github.com/goark/mt/mt19937"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Define method for CardService
type ICardService interface {
	GetCards() ([]schemas.Card, error)
	PickRandomCards([]schemas.Card, int) ([]string, error)
}

// Define struct for CardService
type CardService struct {
	db *gorm.DB
}

// Implement method for CardService struct
func CardServiceImpl() ICardService {
	db := database.DBConnect()
	err := db.AutoMigrate(&schemas.Card{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate database")
	}

	return &CardService{
		db: db,
	}
}

// GetCards is a method for CardService struct
func (service *CardService) GetCards() ([]schemas.Card, error) {
	// Get all cards
	var cards []schemas.Card
	result := service.db.Table("cards").Find(&cards)
	if result.Error != nil {
		log.Err(result.Error).Msg("Failed to get cards")
		return nil, result.Error
	}
	log.Info().Msgf("Get cards successfully: %v", result)
	return cards, nil
}

// PickRandomCards is a method for CardService struct
func (service *CardService) PickRandomCards(cards []schemas.Card, count int) ([]string, error) {
	r := rand.New(mt19937.New(time.Now().UnixNano()))

	// Durstenfeld shuffle
	for i := len(cards) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	// Randomly pick cards from the shuffled list
	finalCards := make([]string, count)
	for i := 0; i < count; i++ {
		// Pick a random index from the cards list
		index := r.Intn(len(cards))
		finalCards[i] = cards[index].CardName
		// Remove the picked card from the cards list to avoid duplicate
		cards = append(cards[:index], cards[index+1:]...)
	}

	return finalCards, nil
}
