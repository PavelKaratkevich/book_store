package domain

type Book struct {
	ID      int `json:"First_name,omitempty" db:"account_id"`
	Title   string
	Authors []string
	Year    string
}
