package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/praveenvoonna/order-matching-service/internal/database"
	"github.com/praveenvoonna/order-matching-service/pkg/models"
)

type ProductHandler struct {
	DB database.Database
}

func NewProductHandler(db database.Database) *ProductHandler {
	return &ProductHandler{DB: db}
}

func (ph *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := ph.DB.QueryRows("SELECT id, name, description, price, quantity FROM order_service.products")
	if err != nil {
		log.Println("Error querying products:", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			log.Println("Error scanning product:", err)
			http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	if len(products) == 0 {
		log.Println("No products found")
		http.Error(w, "No products found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Println("Error encoding products:", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println("Error decoding product:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	statement := `INSERT INTO order_service.products (name, description, price, quantity) VALUES ($1, $2, $3, $4) RETURNING id`
	err = ph.DB.QueryRow(statement, product.Name, product.Description, product.Price, product.Quantity).Scan(&product.ID)
	if err != nil {
		log.Println("Error creating product:", err)
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		log.Println("Error encoding product:", err)
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}
}

func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println("Error decoding product:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	productID := r.URL.Query().Get("id")
	statement := `UPDATE order_service.products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5`
	_, err = ph.DB.Execute(statement, product.Name, product.Description, product.Price, product.Quantity, productID)
	if err != nil {
		log.Println("Error updating product:", err)
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		log.Println("Error encoding product:", err)
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}
}

func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	statement := `DELETE FROM order_service.products WHERE id = $1`
	_, err := ph.DB.Execute(statement, productID)
	if err != nil {
		log.Println("Error deleting product:", err)
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
