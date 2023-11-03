package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/praveenvoonna/order-matching-service/internal/database"
	"github.com/praveenvoonna/order-matching-service/pkg/models"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	rows, err := db.QueryRows("SELECT id, name, description, price, quantity FROM order_service.products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	json.NewEncoder(w).Encode(products)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	statement := `INSERT INTO order_service.products (name, description, price, quantity) VALUES ($1, $2, $3, $4) RETURNING id`
	err = db.QueryRow(statement, product.Name, product.Description, product.Price, product.Quantity).Scan(&product.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productID := r.URL.Query().Get("id")
	statement := `UPDATE order_service.products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5`
	_, err = db.Execute(statement, product.Name, product.Description, product.Price, product.Quantity, productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	productID := r.URL.Query().Get("id")
	statement := `DELETE FROM order_service.products WHERE id = $1`
	_, err := db.Execute(statement, productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
