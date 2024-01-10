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
	Items        []Vacancy   `json:"items"`
	Found        int         `json:"found"`
	Pages        int         `json:"pages"`
	Page         int         `json:"page"`
	PerPage      int         `json:"per_page"`
	Clusters     interface{} `json:"clusters"`
	Arguments    interface{} `json:"arguments"`
	Fixes        interface{} `json:"fixes"`
	Suggests     interface{} `json:"suggests"`
	AlternateUrl string      `json:"alternate_url"`
}

func NewVacancies() Vacancies {
	return Vacancies{}
}
