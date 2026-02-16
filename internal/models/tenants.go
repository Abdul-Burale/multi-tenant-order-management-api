package models

import (
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	TenantID     uuid.UUID `json:"tenant_id"`
	BusinessName string    `json:"business_name"`
	Created_At   time.Time `json:"created_at"`
}
