package models

import (
	"time"
)

type Product struct {
	ID        uint   `json:"id"`
	Name  string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Sku string `json:"sku"`
	ImageUrl string `json:"image_url"`
	Notes string `json:"notes"`
	Price        uint   `json:"price"`
	Stock        uint   `json:"stock"`
	Location        string   `json:"location"`
	IsAvailable bool `json:"is_available"`
}

type CreateProduct struct {
	Name  string `json:"name"`
	Sku string `json:"sku"`
	ImageUrl string `json:"image_url"`
	Notes string `json:"notes"`
	Price        uint   `json:"price"`
	Stock        uint   `json:"stock"`
	Location        string   `json:"location"`
	IsAvailable bool `json:"is_available"`
}
