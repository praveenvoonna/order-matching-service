package main

import (
	"log"
	"net/http"

	"github.com/praveenvoonna/order-matching-service/internal/database"
	"github.com/praveenvoonna/order-matching-service/internal/handlers"
)

func main() {
	db := &database.PostgresDatabase{}

	db.Connect()

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetOrdersHandler(w, r, db)
	})

	http.HandleFunc("/create_order", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateOrderHandler(w, r, db)
	})

	http.HandleFunc("/update_order", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateOrderHandler(w, r, db)
	})

	http.HandleFunc("/delete_order", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteOrderHandler(w, r, db)
	})

	http.HandleFunc("/buyers", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetBuyersHandler(w, r, db)
	})

	http.HandleFunc("/create_buyer", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateBuyerHandler(w, r, db)
	})

	http.HandleFunc("/update_buyer", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateBuyerHandler(w, r, db)
	})

	http.HandleFunc("/delete_buyer", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteBuyerHandler(w, r, db)
	})

	http.HandleFunc("/sellers", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetSellersHandler(w, r, db)
	})

	http.HandleFunc("/create_seller", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateSellerHandler(w, r, db)
	})

	http.HandleFunc("/update_seller", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateSellerHandler(w, r, db)
	})

	http.HandleFunc("/delete_seller", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteSellerHandler(w, r, db)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
