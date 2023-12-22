package model

type Employer struct {
	Name string `json:"name"`
}

type Vacancy struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	PublishedAt  string   `json:"published_at"`
	CreatedAt    string   `json:"created_at"`
	Archived     bool     `json:"archived"`
	AlternateUrl string   `json:"alternate_url"`
	Employer     Employer `json:"employer"`
}
type VacancyList struct {
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

func NewVacancyList() VacancyList {
	return VacancyList{}
}
