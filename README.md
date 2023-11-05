# Order Matching Service

This service provides functionality for managing products, orders, buyers, and sellers. It enables users to perform CRUD operations on these entities through a RESTful API.

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

Please refer to postman collection file under `resources/postman/order-matching-service.postman_collection.json`

## Configuration
The service uses a PostgreSQL database to store and manage data. The connection string for the database can be configured in the `internal/database/database.go` file. Ensure that the PostgreSQL database is properly configured and accessible before running the service.