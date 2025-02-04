package config

type Settings struct {
	VerbosityLevel string
	ProviderName   string
	TemplateID     string
}

var GlobalSettings Settings

func GetSettings() *Settings {
	return &GlobalSettings
}
