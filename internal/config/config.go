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

const (
	SimilarVacancies = "https://api.hh.ru/resumes/%s/similar_vacancies?id=publication_time&page=%d&per_page=100"
	AllVacancies     = "https://api.hh.ru/vacancies?text=golang&id=publication_time&page=%d&per_page=100"
	Resume           = "https://api.hh.ru/resumes/mine"
)
