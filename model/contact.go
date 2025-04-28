package model

type Contact struct {
	ID    int    `json:"contact_id"`
	Name  string `json:"contact_name"`
	Email string `json:"contact_email"`
}