CREATE TABLE IF NOT EXISTS product (
    id VARCHAR(20) PRIMARY KEY,
    name TEXT NOT NULL,
    price DECIMAL(10, 5) NOT NULL
);

CREATE TABLE IF NOT EXISTS price_history (
    id VARCHAR(20) PRIMARY KEY,
    product_id VARCHAR(20) NOT NULL,
    price DECIMAL(10, 5) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_product_id FOREIGN KEY (product_id) REFERENCES product(id)
);