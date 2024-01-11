package database

import (
	"fmt"
	"golang_template/util"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB {
	// Load config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	// Connect to database
	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect database")
	}

	return db
}
