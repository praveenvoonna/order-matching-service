package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/praveenvoonna/order-matching-service/internal/database"
	"github.com/praveenvoonna/order-matching-service/pkg/models"
)

func GetOrdersHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	rows, err := db.QueryRows("SELECT id, buyer_id, seller_id, quantity, price, status FROM order_service.orders")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.BuyerID, &order.SellerID, &order.Quantity, &order.Price, &order.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	json.NewEncoder(w).Encode(orders)
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	statement := `INSERT INTO order_service.orders (buyer_id, seller_id, quantity, price, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = db.QueryRow(statement, order.BuyerID, order.SellerID, order.Quantity, order.Price, order.Status).Scan(&order.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func UpdateOrderHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	statement := `UPDATE order_service.orders SET buyer_id = $1, seller_id = $2, quantity = $3, price = $4, status = $5 WHERE id = $6`
	_, err = db.Execute(statement, order.BuyerID, order.SellerID, order.Quantity, order.Price, order.Status, orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func DeleteOrderHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	orderID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	statement := `DELETE FROM order_service.orders WHERE id = $1`
	_, err = db.Execute(statement, orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
