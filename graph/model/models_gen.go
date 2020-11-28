// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BlogPost struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NewBlogPost struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NewProject struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Language    string `json:"language"`
}

type Project struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Language    string `json:"language"`
	Description string `json:"description"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
