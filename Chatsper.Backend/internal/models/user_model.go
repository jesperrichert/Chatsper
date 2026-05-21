package models

import "gorm.io/gorm"

type ExampleEntity struct {
	gorm.Model
	Name  string
	Email string
}

// DTOs - Data transfer objects
type ExampleRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ExampleResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
