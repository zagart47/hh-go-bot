package entity

type AppliedVacancy struct {
	Overall Overall `json:"overall"`
}

type Overall struct {
	AlreadyApplied int64 `json:"already_applied"`
}

func NewAppliedVacancy() AppliedVacancy {
	return AppliedVacancy{}
}
