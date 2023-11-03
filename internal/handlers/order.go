package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/praveenvoonna/order-matching-service/internal/database"
	"github.com/praveenvoonna/order-matching-service/pkg/models"
)

type OrderHandler struct {
	DB database.Database
}

func NewOrderHandler(db database.Database) *OrderHandler {
	return &OrderHandler{DB: db}
}

func (oh *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	rows, err := oh.DB.QueryRows("SELECT id, buyer_id, seller_id, quantity, price, status FROM order_service.orders")
	if err != nil {
		log.Println("Error querying orders:", err)
		http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.BuyerID, &order.SellerID, &order.Quantity, &order.Price, &order.Status)
		if err != nil {
			log.Println("Error scanning order:", err)
			http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	if len(orders) == 0 {
		log.Println("No orders found")
		http.Error(w, "No orders found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(orders); err != nil {
		log.Println("Error encoding orders:", err)
		http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}
}

func (oh *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println("Error decoding order:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	statement := `INSERT INTO order_service.orders (buyer_id, seller_id, quantity, price, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = oh.DB.QueryRow(statement, order.BuyerID, order.SellerID, order.Quantity, order.Price, order.Status).Scan(&order.ID)
	if err != nil {
		log.Println("Error creating order:", err)
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (oh *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println("Error decoding order:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	orderID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Invalid order ID:", err)
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	statement := `UPDATE order_service.orders SET buyer_id = $1, seller_id = $2, quantity = $3, price = $4, status = $5 WHERE id = $6`
	_, err = oh.DB.Execute(statement, order.BuyerID, order.SellerID, order.Quantity, order.Price, order.Status, orderID)
	if err != nil {
		log.Println("Error updating order:", err)
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func (oh *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Invalid order ID:", err)
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	statement := `DELETE FROM order_service.orders WHERE id = $1`
	_, err = oh.DB.Execute(statement, orderID)
	if err != nil {
		log.Println("Error deleting order:", err)
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
