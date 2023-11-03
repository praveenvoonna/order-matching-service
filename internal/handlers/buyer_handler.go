package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/praveenvoonna/order-matching-service/internal/database"
	"github.com/praveenvoonna/order-matching-service/pkg/models"
)

func GetBuyersHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	rows, err := db.Query("SELECT id, name FROM order_service.buyers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var buyers []models.Buyer
	for rows.Next() {
		var buyer models.Buyer
		err := rows.Scan(&buyer.ID, &buyer.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		buyers = append(buyers, buyer)
	}

	json.NewEncoder(w).Encode(buyers)
}

func CreateBuyerHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	var buyer models.Buyer
	err := json.NewDecoder(r.Body).Decode(&buyer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	statement := `INSERT INTO order_service.buyers (name) VALUES ($1) RETURNING id`
	err = db.QueryRow(statement, buyer.Name).Scan(&buyer.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(buyer)
}

func UpdateBuyerHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	var buyer models.Buyer
	err := json.NewDecoder(r.Body).Decode(&buyer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buyerID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid buyer ID", http.StatusBadRequest)
		return
	}

	statement := `UPDATE order_service.buyers SET name = $1 WHERE id = $2`
	_, err = db.Exec(statement, buyer.Name, buyerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(buyer)
}

func DeleteBuyerHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	buyerID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid buyer ID", http.StatusBadRequest)
		return
	}

	statement := `DELETE FROM order_service.buyers WHERE id = $1`
	_, err = db.Exec(statement, buyerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
