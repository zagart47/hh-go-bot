package config

import "github.com/ilyakaznacheev/cleanenv"

type config struct {
	Token    string `yaml:"token"`
	Bearer   string `yaml:"bearer"`
	ResumeID string `yaml:"resume_id"`
}

func New() (config, error) {
	cfg := config{}
	if err := cleanenv.ReadConfig("./internal/config/config.yaml", &cfg); err != nil {
		return config{}, err
	}
	return cfg, nil
}
