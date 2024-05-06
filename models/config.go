package models

type Config struct {
	Provider struct {
		Name    string `json:"name"`
		Gateway string `json:"gateway"`
	} `json:"provider"`
	Functions []struct {
		Name    string `yaml:"name"`
		Lang    string `yaml:"lang"`
		Handler string `yaml:"handler"`
		Image   string `yaml:"image"`
		Env     struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
		} `yaml:"env"`
	} `yaml:"functions"`
}
