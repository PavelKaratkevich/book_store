package dto

import (
	"github.com/lib/pq"
)

type Book struct {
	ID      int            `json:"ID" db:"id"`
	Title   string         `json:"Title" db:"title"`
	Authors pq.StringArray `json:"Authors" db:"authors"`
	Year    string         `json:"Year" db:"year"`
}

