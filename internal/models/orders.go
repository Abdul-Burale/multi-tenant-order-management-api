package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusDispatched OrderStatus = "dispatched"
	OrderStatusDone       OrderStatus = "done"
)

type Order struct {
	OrderID         uuid.UUID   `json:"order_id"`
	StoreID         uuid.UUID   `json:"store_id"`
	UserID          uuid.UUID   `json:"user_id"`
	Status          OrderStatus `json:"status"`
	TotalAmount     int         `json:"total_amount"`
	StripePaymentID *string     `json:"stripe_payment_id,omitempty"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}
