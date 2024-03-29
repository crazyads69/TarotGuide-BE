package util

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	_ "golang_template/docs"
)

// Config stores all application configs loaded from file or env variables
type Config struct {
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`

	SwaggerURL    string `mapstructure:"SWAGGER_URL"`
	DBPath        string `mapstructure:"DB_PATH"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        int    `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	LogFilename   string `mapstructure:"LOG_FILENAME"`
	LogMaxSize    int    `mapstructure:"LOG_MAX_SIZE"` // in megabytes
	LogMaxBackups int    `mapstructure:"LOG_MAX_BACKUPS"`
	LogMaxAge     int    `mapstructure:"LOG_MAX_AGE"`
	LogCompress   bool   `mapstructure:"LOG_COMPRESS"`
	GeminiKey     string `mapstructure:"GEMINI_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("HTTP_SERVER_ADDRESS", ":8080")
	viper.SetDefault("SWAGGER_URL", "/docs")
	viper.SetDefault("LOG_FILENAME", "log/app.log")
	viper.SetDefault("LOG_MAX_SIZE", 10)
	viper.SetDefault("LOG_MAX_BACKUPS", 5)
	viper.SetDefault("LOG_MAX_AGE", 28)
	viper.SetDefault("LOG_COMPRESS", true)
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("DB_NAME", "DB")
	viper.SetDefault("DB_PATH", "./database/app.db")
	viper.SetDefault("GEMINI_KEY", "")

	// Read config
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Unmarshal config
	err = viper.Unmarshal(&config)
	return
}

func WatchConfig(config *Config) {
	// Watch config file changes
	viper.WatchConfig()
	// Handle config file changes
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		// Unmarshal config file
		err := viper.Unmarshal(&config)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to unmarshal config")
		}
		// Reload config file
		LoadConfig(".")
	})
}
