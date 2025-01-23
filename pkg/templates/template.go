package templates

type Template struct {
	ID        string                    `yaml:"id"`
	Info      Info                      `yaml:"info"`
	Providers map[string]ProviderConfig `yaml:"providers"`
}

type Info struct {
	Name         string                 `yaml:"name"`
	Author       string                 `yaml:"author"`
	Description  string                 `yaml:"description"`
	References   []string               `yaml:"references"`
	Technologies []string               `yaml:"technologies"`
	Tags         []string               `yaml:"tags"`
	Metadata     map[string]interface{} `yaml:"metadata"`
}

type ProviderConfig struct {
	Targets  []string `yaml:"targets,omitempty"`
	Content  string   `yaml:"content,omitempty"`
	Setup    string   `yaml:"setup,omitempty"`
	Teardown string   `yaml:"teardown,omitempty"`
}
