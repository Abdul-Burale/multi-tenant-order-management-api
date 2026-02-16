package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderItems struct {
	OrderItemsID uuid.UUID `json:"order_items_id"`
	OrderID      uuid.UUID `json:"order_id"`
	ProductID    uuid.UUID `json:"product_id"`
	Quantity     int       `json:"quantity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
