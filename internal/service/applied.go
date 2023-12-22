package service

import (
	"hh-go-bot/internal/model"
)

type Vacancies interface {
	ShowAppliedVacancies() model.VacancyList
}
