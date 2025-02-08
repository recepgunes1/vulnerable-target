package logger

import (
	"os"
	"time"

	"github.com/happyhackingspace/vulnerable-target/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	settings := config.GetSettings()

	level, err := zerolog.ParseLevel(settings.VerbosityLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}

	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.TimeOnly,
		NoColor:    false,
	}

	log.Logger = zerolog.
		New(output).
		Level(level).
		With().
		Timestamp().
		Logger()
}
