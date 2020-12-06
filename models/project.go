package models

import "time"

type Project struct {
	ID    int    `json:"id" pg:",pk,unique,notnull"`
	Title string `json:"title"`

	Description string `json:"description"`
	Language  string `json:"language"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}


