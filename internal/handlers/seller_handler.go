package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/praveenvoonna/order-matching-service/internal/database"
	"github.com/praveenvoonna/order-matching-service/pkg/models"
)

func GetSellersHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	rows, err := db.QueryRows("SELECT id, name FROM order_service.sellers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sellers []models.Seller
	for rows.Next() {
		var seller models.Seller
		err := rows.Scan(&seller.ID, &seller.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sellers = append(sellers, seller)
	}

	json.NewEncoder(w).Encode(sellers)
}

func CreateSellerHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	var seller models.Seller
	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	statement := `INSERT INTO order_service.sellers (name) VALUES ($1) RETURNING id`
	err = db.QueryRow(statement, seller.Name).Scan(&seller.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(seller)
}

func UpdateSellerHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	var seller models.Seller
	err := json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sellerID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	statement := `UPDATE order_service.sellers SET name = $1 WHERE id = $2`
	_, err = db.Execute(statement, seller.Name, sellerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(seller)
}

func DeleteSellerHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	sellerID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid seller ID", http.StatusBadRequest)
		return
	}

	statement := `DELETE FROM order_service.sellers WHERE id = $1`
	_, err = db.Execute(statement, sellerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
