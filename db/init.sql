CREATE SCHEMA IF NOT EXISTS order_service;

CREATE TABLE IF NOT EXISTS order_service.products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC NOT NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_service.buyers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    budget NUMERIC NOT NULL,
    quantity INTEGER NOT NULL,
    product_id INTEGER REFERENCES order_service.products(id)
);

CREATE TABLE IF NOT EXISTS order_service.sellers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    price NUMERIC NOT NULL,
    product VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS order_service.orders (
    id SERIAL PRIMARY KEY,
    buyer_id INTEGER REFERENCES order_service.buyers(id),
    seller_id INTEGER REFERENCES order_service.sellers(id),
    quantity INTEGER NOT NULL,
    price NUMERIC NOT NULL,
    status VARCHAR(50) NOT NULL
);
