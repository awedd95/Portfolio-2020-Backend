package models
import "time"

type BlogPost struct {
	ID    int    `json:"id" pg:",pk,unique,notnull"`
    Title    string `json:"title"`
    Body    string `json:"body"`
	
    CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
