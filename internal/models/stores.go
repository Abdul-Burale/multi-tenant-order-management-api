package models

import (
	"time"

	"github.com/google/uuid"
)

type StoreType string

const (
	StoreTypeOnline   StoreType = "online"
	StoreTypePhyiscal StoreType = "physical"
)

type Stores struct {
	StoreID     uuid.UUID `json:"store_id"`
	TenantID    uuid.UUID `json:"tenant_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Type        StoreType `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
}
