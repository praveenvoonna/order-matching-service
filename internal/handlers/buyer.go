package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/praveenvoonna/order-matching-service/internal/database"
	"github.com/praveenvoonna/order-matching-service/pkg/models"
)

type BuyerHandler struct {
	DB database.Database
}

func NewBuyerHandler(db database.Database) *BuyerHandler {
	return &BuyerHandler{DB: db}
}

func (bh *BuyerHandler) GetBuyers(w http.ResponseWriter, r *http.Request) {
	rows, err := bh.DB.QueryRows("SELECT id, name, email FROM order_service.buyers")
	if err != nil {
		log.Println("Error querying buyers:", err)
		http.Error(w, "Failed to fetch buyers", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var buyers []models.Buyer
	for rows.Next() {
		var buyer models.Buyer
		err := rows.Scan(&buyer.ID, &buyer.Name, &buyer.Email)
		if err != nil {
			log.Println("Error scanning buyer:", err)
			http.Error(w, "Failed to fetch buyers", http.StatusInternalServerError)
			return
		}
		buyers = append(buyers, buyer)
	}

	if len(buyers) == 0 {
		log.Println("No buyers found")
		http.Error(w, "No buyers found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(buyers); err != nil {
		log.Println("Error encoding buyers:", err)
		http.Error(w, "Failed to fetch buyers", http.StatusInternalServerError)
		return
	}
}

func (bh *BuyerHandler) CreateBuyer(w http.ResponseWriter, r *http.Request) {
	var buyer models.Buyer
	err := json.NewDecoder(r.Body).Decode(&buyer)
	if err != nil {
		log.Println("Error decoding buyer:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	statement := `INSERT INTO order_service.buyers (name, email, product_id) VALUES ($1, $2, $3) RETURNING id`
	err = bh.DB.QueryRow(statement, buyer.Name, buyer.Email, buyer.ProductID).Scan(&buyer.ID)
	if err != nil {
		log.Println("Error creating buyer:", err)
		http.Error(w, "Failed to create buyer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(buyer)
}

func (bh *BuyerHandler) UpdateBuyer(w http.ResponseWriter, r *http.Request) {
	var buyer models.Buyer
	err := json.NewDecoder(r.Body).Decode(&buyer)
	if err != nil {
		log.Println("Error decoding buyer:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	buyerID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Invalid buyer ID:", err)
		http.Error(w, "Invalid buyer ID", http.StatusBadRequest)
		return
	}

	statement := `UPDATE order_service.buyers SET name = $1, email = $2, product_id = $3 WHERE id = $4`
	_, err = bh.DB.Execute(statement, buyer.Name, buyer.Email, buyer.ProductID, buyerID)
	if err != nil {
		log.Println("Error updating buyer:", err)
		http.Error(w, "Failed to update buyer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(buyer)
}

func (bh *BuyerHandler) DeleteBuyer(w http.ResponseWriter, r *http.Request) {
	buyerID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Invalid buyer ID:", err)
		http.Error(w, "Invalid buyer ID", http.StatusBadRequest)
		return
	}

	statement := `DELETE FROM order_service.buyers WHERE id = $1`
	_, err = bh.DB.Execute(statement, buyerID)
	if err != nil {
		log.Println("Error deleting buyer:", err)
		http.Error(w, "Failed to delete buyer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
