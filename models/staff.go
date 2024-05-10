package models

import (
	"time"
)

type Staff struct {
	ID        uint   `json:"id"`
	Name  string `json:"name"`
	Password string `json:"password"`
	PhoneNumber  string `json:"phone_number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InputCreateStaff struct {
	Name  string `json:"name"`
	Password string `json:"password"`
	PhoneNumber  string `json:"phone_number"`
}

type InputLoginStaff struct {
	Name  string `json:"name"`
	Password string `json:"password"`
}