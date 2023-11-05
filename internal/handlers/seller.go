package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/praveenvoonna/order-matching-service/internal/database"
	"github.com/praveenvoonna/order-matching-service/pkg/models"
)

type SellerHandler struct {
	DB database.Database
}

func NewSellerHandler(db database.Database) *SellerHandler {
	return &SellerHandler{DB: db}
}

func (sh *SellerHandler) GetSellers(w http.ResponseWriter, r *http.Request) {
	rows, err := sh.DB.QueryRows("SELECT id, name, email, product_id FROM order_service.sellers")
	if err != nil {
		log.Println("Error querying sellers:", err)
		http.Error(w, "Failed to fetch sellers", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sellers []models.Seller
	for rows.Next() {
		var seller models.Seller
		err := rows.Scan(&seller.ID, &seller.Name, &seller.Email, &seller.ProductID)
		if err != nil {
			log.Println("Error scanning seller:", err)
			http.Error(w, "Failed to fetch sellers", http.StatusInternalServerError)
			return
		}
		sellers = append(sellers, seller)
	}

	if len(sellers) == 0 {
		log.Println("No sellers found")
		http.Error(w, "No sellers found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(sellers); err != nil {
		log.Println("Error encoding sellers:", err)
		http.Error(w, "Failed to fetch sellers", http.StatusInternalServerError)
		return
	}
}

func (sh *SellerHandler) CreateSeller(w http.ResponseWriter, r *http.Request) {
	var seller models.Seller
	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		log.Println("Error decoding seller:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	statement := `INSERT INTO order_service.sellers (name, email, product_id) VALUES ($1, $2, $3) RETURNING id`
	err = sh.DB.QueryRow(statement, seller.Name, seller.Email, seller.ProductID).Scan(&seller.ID)
	if err != nil {
		log.Println("Error creating seller:", err)
		http.Error(w, "Failed to create seller", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(seller)
}

func (sh *SellerHandler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
	var seller models.Seller
	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		log.Println("Error decoding seller:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	sellerID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Invalid seller ID:", err)
		http.Error(w, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	statement := `UPDATE order_service.sellers SET name = $1, email = $2, product_id = $3 WHERE id = $4`
	_, err = sh.DB.Execute(statement, seller.Name, seller.Email, seller.ProductID, sellerID)
	if err != nil {
		log.Println("Error updating seller:", err)
		http.Error(w, "Failed to update seller", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(seller)
}

func (sh *SellerHandler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	sellerID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Invalid seller ID:", err)
		http.Error(w, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	statement := `DELETE FROM order_service.sellers WHERE id = $1`
	_, err = sh.DB.Execute(statement, sellerID)
	if err != nil {
		log.Println("Error deleting seller:", err)
		http.Error(w, "Failed to delete seller", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
