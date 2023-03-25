package responsemodel

// StaffAll ...
type StaffAll struct {
	List  []string `json:"list"`
	Page  int64    `json:"page"`
	Limit int64    `json:"limit"`
}

// StaffBrief ...
type StaffBrief struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}
