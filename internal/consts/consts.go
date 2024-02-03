package consts

import "time"

const (
	BOT                   = "bot"
	HTTP                  = "http"
	Timeout time.Duration = 9999999

	AllVacanciesLink     = "https://api.hh.ru/vacancies?text=golang&area=113&order_by=publication_time&search_field=name&per_page=100&page=%d"
	SimilarVacanciesLink = "https://api.hh.ru/resumes/%s/similar_vacancies?per_page=%d&page=%d&id=publication_time"
	ResumeLink           = "https://api.hh.ru/resumes/mine"

	NoExp        = "noExperience"
	Between1and3 = "between1And3"
	Between3and6 = "between3And6"
	MoreThan6    = "moreThan6"
)
