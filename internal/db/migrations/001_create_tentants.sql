CREATE TABLE tenants (
    tenantID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
