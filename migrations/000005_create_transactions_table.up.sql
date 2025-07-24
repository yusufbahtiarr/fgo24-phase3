CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    transaction_type transaction_type_enum NOT NULL,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE,
    quantity INT NOT NULL,
    unit_price DECIMAL(12,2) NOT NULL,
    total_price DECIMAL(12,2) NOT NULL,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);