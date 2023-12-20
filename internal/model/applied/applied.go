package applied

type Vacancy struct {
	Items   []Item  `json:"items"`
	Found   int64   `json:"found"`
	Pages   int64   `json:"pages"`
	Page    int64   `json:"page"`
	PerPage int64   `json:"per_page"`
	Overall Overall `json:"overall"`
}
type VacancyList []Vacancy

type Item struct {
	LastName           string          `json:"last_name"`
	FirstName          string          `json:"first_name"`
	MiddleName         interface{}     `json:"middle_name"`
	Title              string          `json:"title"`
	CreatedAt          string          `json:"created_at"`
	UpdatedAt          string          `json:"updated_at"`
	Area               Area            `json:"area"`
	Age                interface{}     `json:"age"`
	Gender             Gender          `json:"gender"`
	Salary             interface{}     `json:"salary"`
	Photo              interface{}     `json:"photo"`
	TotalExperience    TotalExperience `json:"total_experience"`
	Certificate        []Certificate   `json:"certificate"`
	HiddenFields       []interface{}   `json:"hidden_fields"`
	Actions            Actions         `json:"actions"`
	URL                string          `json:"url"`
	AlternateURL       string          `json:"alternate_url"`
	ID                 string          `json:"id"`
	Download           Download        `json:"download"`
	Platform           Platform        `json:"platform"`
	Education          Education       `json:"education"`
	Experience         []Experience    `json:"experience"`
	Marked             bool            `json:"marked"`
	Finished           bool            `json:"finished"`
	Status             Gender          `json:"status"`
	Access             Access          `json:"access"`
	RequiresCompletion bool            `json:"requires_completion"`
	HasErrors          bool            `json:"has_errors"`
}

type Access struct {
	Type Gender `json:"type"`
}

type Gender struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Actions struct {
	Download Download `json:"download"`
}

type Download struct {
	PDF PDF `json:"pdf"`
	Rtf PDF `json:"rtf"`
}

type PDF struct {
	URL string `json:"url"`
}

type Area struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Certificate struct {
	Owner      interface{} `json:"owner"`
	Type       string      `json:"type"`
	Title      string      `json:"title"`
	AchievedAt string      `json:"achieved_at"`
	URL        string      `json:"url"`
}

type Education struct {
	Level   Gender    `json:"level"`
	Primary []Primary `json:"primary"`
}

type Primary struct {
	Name           string      `json:"name"`
	Organization   string      `json:"organization"`
	Result         string      `json:"result"`
	Year           int64       `json:"year"`
	NameID         string      `json:"name_id"`
	OrganizationID interface{} `json:"organization_id"`
	ResultID       string      `json:"result_id"`
}

type Experience struct {
	Start      string      `json:"start"`
	End        *string     `json:"end"`
	Company    string      `json:"company"`
	CompanyID  *string     `json:"company_id"`
	Industry   interface{} `json:"industry"`
	Industries []Gender    `json:"industries"`
	Area       Area        `json:"area"`
	CompanyURL interface{} `json:"company_url"`
	Employer   interface{} `json:"employer"`
	Position   *string     `json:"position,omitempty"`
}

type Platform struct {
	ID string `json:"id"`
}

type TotalExperience struct {
	Months int64 `json:"months"`
}

type Overall struct {
	NotPublished   int64 `json:"not_published"`
	AlreadyApplied int64 `json:"already_applied"`
	Unavailable    int64 `json:"unavailable"`
}
