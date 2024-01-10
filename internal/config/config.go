package config

import "github.com/ilyakaznacheev/cleanenv"

type Cfg struct {
	Bot struct {
		Token string `yaml:"token"`
	} `yaml:"bot"`
	HTTP struct {
		Host string `yaml:"host"`
	} `yaml:"http"`
	Api struct {
		Bearer   string `yaml:"bearer"`
		ResumeID string `yaml:"resume_id"`
	} `yaml:"api"`
}

func All() (Cfg, error) {
	cfg := Cfg{}
	if err := cleanenv.ReadConfig("./internal/config/config.yaml", &cfg); err != nil {
		return Cfg{}, err
	}
	return cfg, nil
}
