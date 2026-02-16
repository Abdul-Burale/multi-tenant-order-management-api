CREATE TABLE stores (
    storeID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL UNIQUE REFERENCES tenants(tenantID) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT,
    type TEXT NOT NULL CHECK (type IN ('online','physical')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);