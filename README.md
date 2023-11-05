# Order Matching Service

This service provides functionality for managing products, orders, buyers, and sellers. It enables users to perform CRUD operations on these entities through a RESTful API.

## Prerequisites

Before running the Order Matching Service, ensure that you have the following installed:

- Go 1.16 or higher
- PostgreSQL

## Installation

To use this service, you can clone the repository and then build and run the main.go file:

```bash
git clone https://github.com/praveenvoonna/order-matching-service.git
cd order-matching-service
go build main.go
./main
```

Ensure that the necessary dependencies are installed using Go modules.

## Usage

This service exposes endpoints for managing products, orders, buyers, and sellers. Each endpoint supports various HTTP methods for performing different actions. Here are the available endpoints:

- `/products`: Manages product data.
- `/orders`: Manages order data.
- `/buyers`: Manages buyer data.
- `/sellers`: Manages seller data.

To interact with the service, you can use tools like cURL or an API client. Here's an example using cURL to get the list of products:

```bash
curl -X GET http://localhost:8080/products

```

You can replace the method and URL to perform different operations on different entities.

Please refer to the postman collection file under `resources/postman/order-matching-service.postman_collection.json` for detailed API documentation.

## Configuration
The service uses a PostgreSQL database to store and manage data. The connection string for the database can be configured in the `internal/database/database.go` file. Ensure that the PostgreSQL database is properly configured and accessible before running the service.

## Schema Setup
To set up the required database schema for the service, you can use the following SQL commands:

```sql
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
    product_id INTEGER REFERENCES order_service.products(id)
);

CREATE TABLE IF NOT EXISTS order_service.sellers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    product_id INTEGER REFERENCES order_service.products(id)
);

CREATE TABLE IF NOT EXISTS order_service.orders (
    id SERIAL PRIMARY KEY,
    buyer_id INTEGER REFERENCES order_service.buyers(id),
    seller_id INTEGER REFERENCES order_service.sellers(id),
    quantity INTEGER NOT NULL,
    price NUMERIC NOT NULL,
    status VARCHAR(50) NOT NULL
);
```

Make sure you have the necessary privileges to create and modify tables within your PostgreSQL database.


Feel free to adjust the provided template to fit the specific prerequisites and instructions required for your Order Matching Service project.
