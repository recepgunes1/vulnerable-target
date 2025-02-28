package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Settings struct {
	VerbosityLevel string
	ProviderName   string
	TemplateID     string
}

var GlobalSettings Settings

func GetSettings() *Settings {
	return &GlobalSettings
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("Error loading .env file: %v", err)
	}
}
