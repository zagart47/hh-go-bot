package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Token  string `yaml:"token"`
	Bearer string `yaml:"bearer"`
}

func New() Config {
	cfg := Config{}
	if err := cleanenv.ReadConfig("./config/config.yaml", &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}
	return cfg
}
