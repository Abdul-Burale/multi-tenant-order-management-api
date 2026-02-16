package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductStatus string

const (
	ProductStatusActive     ProductStatus = "active"
	ProductStatusOutOfStock ProductStatus = "out_of_stock"
)

type Products struct {
	ProductID   uuid.UUID     `json:"product_id"`
	StoreID     uuid.UUID     `json:"store_id"`
	Name        string        `json:"name"`
	Description *string       `json:"description,omitempty"`
	Price       int           `json:"price"`
	Status      ProductStatus `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
