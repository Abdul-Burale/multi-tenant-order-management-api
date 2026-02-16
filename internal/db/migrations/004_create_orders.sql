CREATE TABLE orders (
    orderID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    store_id UUID NOT NULL REFERENCES stores(storeID) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(userID) ON DELETE CASCADE,
    status TEXT NOT NULL CHECK (status IN ('pending','dispatched','done')),
    total_amount INT NOT NULL,
    stripe_payment_id TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
