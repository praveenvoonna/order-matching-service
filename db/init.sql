CREATE SCHEMA order_service;

CREATE TABLE order_service.buyers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    budget NUMERIC NOT NULL,
    quantity INTEGER NOT NULL,
    product_id INTEGER REFERENCES products(id)
);

CREATE TABLE order_service.sellers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    price NUMERIC NOT NULL,
    product VARCHAR(255) NOT NULL
);

CREATE TABLE order_service.orders (
    id SERIAL PRIMARY KEY,
    buyer_id INTEGER REFERENCES buyers(id),
    seller_id INTEGER REFERENCES sellers(id),
    quantity INTEGER NOT NULL,
    price NUMERIC NOT NULL,
    status VARCHAR(50) NOT NULL
);

