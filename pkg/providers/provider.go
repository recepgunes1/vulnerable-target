package providers

import (
	"github.com/happyhackingspace/vulnerable-target/internal/config"
	"github.com/happyhackingspace/vulnerable-target/pkg/providers/docker"
	"github.com/rs/zerolog/log"
)

func Start() {
	settings := config.GetSettings()
	switch settings.ProviderName {
	case "docker":
		docker.Run()
	case "azure":
		log.Info().Msgf("azure is not implmented")
	}
	log.Info().Msgf("%s template is running on %s", settings.TemplateID, settings.ProviderName)
}
