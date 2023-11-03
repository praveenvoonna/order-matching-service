package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/praveenvoonna/order-matching-service/internal/database"
	"github.com/praveenvoonna/order-matching-service/pkg/models"
)

func GetOrdersHandler(w http.ResponseWriter, r *http.Request, db database.Database) {
	rows, err := db.Query("SELECT id, product, quantity, price FROM order_service.orders")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.Product, &order.Quantity, &order.Price)
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

	statement := `INSERT INTO order_service.orders (product, quantity, price) VALUES ($1, $2, $3) RETURNING id`
	err = db.QueryRow(statement, order.Product, order.Quantity, order.Price).Scan(&order.ID)
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

	statement := `UPDATE order_service.orders SET product = $1, quantity = $2, price = $3 WHERE id = $4`
	_, err = db.Exec(statement, order.Product, order.Quantity, order.Price, orderID)
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
	_, err = db.Exec(statement, orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
