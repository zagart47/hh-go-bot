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
	PostgreSQL struct {
		Host     string `yaml:"host" default:"localhost"`
		Port     string `yaml:"port" default:"5432"`
		DBName   string `yaml:"db_name" default:"postgres"`
		UserName string `yaml:"user_name" default:"postgres"`
		Pwd      string `yaml:"pwd" default:"postgres"`
	} `yaml:"postgreSQL"`
	Redis struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Pwd     string `yaml:"pwd"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"redis"`
	LoggerMode string `yaml:"logger_mode"`
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
