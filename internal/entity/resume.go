package entity

type Resume struct {
	Items   []Item `json:"items"`
	Found   int64  `json:"found"`
	Pages   int64  `json:"pages"`
	Page    int64  `json:"page"`
	PerPage int64  `json:"per_page"`
}

func NewResume() Resume {
	return Resume{}
}

type Item struct {
	ID string `json:"id"`
}
