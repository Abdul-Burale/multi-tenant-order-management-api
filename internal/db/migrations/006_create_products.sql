CREATE TABLE products (
    productID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES stores(storeID) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT,
    price INT NOT NULL,  -- store as cents
    status TEXT NOT NULL CHECK (status IN ('active','out_of_stock')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
