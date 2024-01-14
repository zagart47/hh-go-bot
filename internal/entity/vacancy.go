package entity

type Employer struct {
	Name string `json:"name"`
}

type Vacancy struct {
	Icon         rune
	Id           string     `json:"id"`
	Name         string     `json:"name"`
	PublishedAt  string     `json:"published_at"`
	CreatedAt    string     `json:"created_at"`
	Archived     bool       `json:"archived"`
	AlternateUrl string     `json:"alternate_url"`
	Relations    []string   `json:"relations"`
	Employer     Employer   `json:"employer"`
	Experience   Experience `json:"experience"`
}
type Experience struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Vacancies struct {
	Items        []Vacancy `json:"items"`
	Pages        int       `json:"pages"`
	AlternateUrl string    `json:"alternate_url"`
}

func NewVacancies() Vacancies {
	return Vacancies{}
}
