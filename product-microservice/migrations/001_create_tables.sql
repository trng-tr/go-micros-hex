-- Active: 1768502153680@@127.0.0.1@5434@goapp2db
CREATE TABLE IF NOT EXISTS products (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    sku VARCHAR(20) UNIQUE NOT NULL,
    category VARCHAR(4) NOT NULL CHECK(category IN('BOOK','CLTH','SHOE')),
    product_name VARCHAR(100) NOT NULL,
    description VARCHAR(255) NOT NULL,
    unit_price BIGINT NOT NULL CHECK (unit_price > 0),
    currency VARCHAR(3) NOT NULL CHECK(currency IN('EUR','USD')),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    is_active BOOLEAN DEFAULT true
 )

 CREATE TABLE IF NOT EXISTS stocks(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    product_id BIGINT NOT NULL UNIQUE,
    quantity BIGINT NOT NULL CHECK (quantity >= 0),
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_stock FOREIGN KEY(product_id) REFERENCES products(id)
 )