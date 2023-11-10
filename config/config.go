package config

import (
	"log"
	"os"

	"github.com/GDGVIT/attendance-app-backend/infra/logger"
	"github.com/spf13/viper"
)

type Configuration struct {
	Server ServerConfiguration
}

// SetupConfig configuration
func SetupConfig() error {
	var configuration *Configuration

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Println("Warning: .env file not found. Using environment variables or defaults.")
	} else {
		// Read .env file
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading .env file: %s", err)
		}
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Errorf("error to decode, %v", err)
		return err
	}

	return nil
}
