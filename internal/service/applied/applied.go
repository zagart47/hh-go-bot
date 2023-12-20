package applied

import (
	"hh-go-bot/internal/model/applied"
)

type Vacancies interface {
	ShowAppliedVacancies() applied.VacancyList
}
