package entity

type Employer struct {
	Name string `json:"name"`
}

type Vacancy struct {
	Icon         rune
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	PublishedAt  string      `json:"published_at"`
	CreatedAt    string      `json:"created_at"`
	Archived     bool        `json:"archived"`
	AlternateUrl string      `json:"alternate_url"`
	Relations    []string    `json:"relations"`
	Employer     Employer    `json:"employer"`
	Exp          Experience  `json:"experience"`
	KeySkills    []KeySkills `json:"key_skills"`
	Description  string      `json:"description"`
}

type KeySkills struct {
	Name string `json:"name"`
}

type Experience struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Vacancies struct {
	Items        []Vacancy `json:"items"`
	Page         int       `json:"page"`
	Pages        int       `json:"pages"`
	AlternateUrl string    `json:"alternate_url"`
	Found        int       `json:"found"`
}

func NewVacancies() Vacancies {
	return Vacancies{}
}

func NewVacancy() Vacancy {
	return Vacancy{}
}
