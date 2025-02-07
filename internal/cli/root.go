package cli

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/happyhackingspace/vulnerable-target/internal/config"
	"github.com/happyhackingspace/vulnerable-target/internal/logger"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var ValidLogLevels = map[string]bool{
	zerolog.DebugLevel.String(): true,
	zerolog.InfoLevel.String():  true,
	zerolog.WarnLevel.String():  true,
	zerolog.ErrorLevel.String(): true,
	zerolog.FatalLevel.String(): true,
	zerolog.PanicLevel.String(): true,
}

var ValidProviders = map[string]bool{
	"aws":           true,
	"azure":         true,
	"google-cloud":  true,
	"digital-ocean": true,
	"docker":        true,
}

func init() {
	settings := config.GetSettings()

	rootCmd.Flags().BoolP("version", "V", false, "Show the current version of the tool")

	rootCmd.Flags().StringVarP(&settings.VerbosityLevel, "verbosity", "v", zerolog.InfoLevel.String(),
		fmt.Sprintf("Set the verbosity level for logs (%s)",
			strings.Join(slices.Collect(maps.Keys(ValidLogLevels)), ", ")))

	rootCmd.Flags().BoolP("list-templates", "l", false, "List all available templates with descriptions")

	rootCmd.Flags().StringVarP(&settings.ProviderName, "provider", "p", "",
		fmt.Sprintf("Specify the cloud provider for building a vulnerable environment (%s)",
			strings.Join(slices.Collect(maps.Keys(ValidProviders)), ", ")))

	rootCmd.Flags().StringVar(&settings.TemplateID, "id", "",
		"Specify a template ID for targeted vulnerable environment")
}

var rootCmd = &cobra.Command{
	Use:     "vt",
	Short:   "Create vulnerable environment",
	Version: "1.0.0",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
		settings := config.GetSettings()

		if len(args) == 0 && cmd.Flags().NFlag() == 0 {
			cmd.Help()
			return
		}

		if versionFlag, _ := cmd.Flags().GetBool("version"); versionFlag {
			fmt.Println(cmd.Version)
			return
		}

		if listTemplates, _ := cmd.Flags().GetBool("list-templates"); listTemplates {
			log.Info().Msg("list templates")
			return
		}

		if !ValidLogLevels[settings.VerbosityLevel] {
			log.Fatal().Msgf("invalid provider '%s'. Valid providers are: %s",
				settings.VerbosityLevel,
				strings.Join(slices.Collect(maps.Keys(ValidLogLevels)), ", "))
		}

		if settings.ProviderName == "" {
			log.Fatal().Msgf("provider is required")
		}

		if !ValidProviders[settings.ProviderName] {
			log.Fatal().Msgf("invalid provider '%s'. Valid providers are: %s",
				settings.ProviderName,
				strings.Join(slices.Collect(maps.Keys(ValidProviders)), ", "))
		}

		if settings.TemplateID == "" {
			log.Fatal().Msgf("template is required")
		}

		log.Info().Msgf("running template %s on %s", settings.TemplateID, settings.ProviderName)
	},
	SilenceErrors: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
