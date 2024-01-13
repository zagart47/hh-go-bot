package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Cfg struct {
	Mode string
	Bot  struct {
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

func NewConfig() Cfg {
	cfg := Cfg{}
	if err := cleanenv.ReadConfig("./internal/config/config.yaml", &cfg); err != nil {
		log.Println("cannot read configs")
		os.Exit(1)
	}
	return cfg
}

var All = NewConfig()

func (c *Cfg) SetMode(m string) {
	c.Mode = m
}
