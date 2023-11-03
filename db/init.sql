CREATE SCHEMA order_service;

CREATE TABLE order_service.orders (
    id SERIAL PRIMARY KEY,
    product TEXT,
    quantity INT,
    price FLOAT
);

CREATE TABLE order_service.sellers (
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE order_service.buyers (
    id SERIAL PRIMARY KEY,
    name TEXT
);
