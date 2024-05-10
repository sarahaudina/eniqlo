package models

import (
	"time"
)

type Checkout struct {
	ID        uint   `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CustomerID  string        `json:"customerId"`
	ProductDetails []ProductDetail `json:"productDetails"`
	Paid          uint        `json:"paid"`
	Change        uint        `json:"change"`
}
  
type ProductDetail struct {
	ProductID string `json:"productId"`
	Quantity  uint   `json:"quantity"`
}

type CheckoutInput struct {
	CustomerID  string        `json:"customerId"`
	ProductDetails ProductDetail `json:"productDetails"`
	Paid          uint        `json:"paid"`
	Change        uint        `json:"change"`
}