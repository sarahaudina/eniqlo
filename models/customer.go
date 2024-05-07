package models

import (
	"time"
)

type Customer struct {
	ID        uint   `json:"id"`
	Name  string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCustomer struct {
	Name  string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
}