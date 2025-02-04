package cli

import (
	"fmt"
	"log"
	"maps"
	"slices"
	"strings"

	"github.com/happyhackingspace/vulnerable-target/internal/config"
	"github.com/spf13/cobra"
)

var ValidLogLevels = map[string]bool{
	"info":  true,
	"warn":  true,
	"error": true,
	"fatal": true,
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

	rootCmd.Flags().StringVarP(&settings.VerbosityLevel, "verbosity", "v", "info",
		fmt.Sprintf("Set the verbosity level for logs (%s)",
			strings.Join(slices.Collect(maps.Keys(ValidLogLevels)), ", ")))

	rootCmd.Flags().BoolP("list-templates", "l", false, "List all available templates with descriptions")

	rootCmd.Flags().StringVarP(&settings.ProviderName, "provider", "p", "",
		fmt.Sprintf("Specify the cloud provider for building a vulnerable environment (%s)",
			strings.Join(slices.Collect(maps.Keys(ValidProviders)), ", ")))

	rootCmd.Flags().StringVar(&settings.TemplateID, "id", "",
		"Specify a template ID for targeted vulnerable environment")

	rootCmd.MarkFlagRequired("provider")

	rootCmd.MarkFlagRequired("id")
}

var rootCmd = &cobra.Command{
	Use:     "vt",
	Short:   "Create vulnerable environment",
	Version: "1.0.0",
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
			log.Println("list templates")
			return
		}

		if !ValidLogLevels[settings.VerbosityLevel] {
			log.Fatalf("invalid provider '%s'. Valid providers are: %s\n",
				settings.VerbosityLevel,
				strings.Join(slices.Collect(maps.Keys(ValidLogLevels)), ", "))
		}

		if !ValidProviders[settings.ProviderName] {
			log.Fatalf("invalid provider '%s'. Valid providers are: %s\n",
				settings.ProviderName,
				strings.Join(slices.Collect(maps.Keys(ValidProviders)), ", "))
		}

		if settings.TemplateID == "" {
			log.Fatalf("template is required")
		}

		log.Printf("running template %s on %s\n", settings.TemplateID, settings.ProviderName)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
