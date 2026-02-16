CREATE TABLE order_items (
    orderItemID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL REFERENCES orders(orderID) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES products(productID) ON DELETE CASCADE,
    quantity INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
