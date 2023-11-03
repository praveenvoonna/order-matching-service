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

	productHandler := handlers.NewProductHandler(db)
	orderHandler := handlers.NewOrderHandler(db)
	buyerHandler := handlers.NewBuyerHandler(db)
	sellerHandler := handlers.NewSellerHandler(db)

	router := http.NewServeMux()

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			productHandler.GetProducts(w, r)
		case "POST":
			productHandler.CreateProduct(w, r)
		case "PUT":
			productHandler.UpdateProduct(w, r)
		case "DELETE":
			productHandler.DeleteProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			orderHandler.GetOrders(w, r)
		case "POST":
			orderHandler.CreateOrder(w, r)
		case "PUT":
			orderHandler.UpdateOrder(w, r)
		case "DELETE":
			orderHandler.DeleteOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/buyers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			buyerHandler.GetBuyers(w, r)
		case "POST":
			buyerHandler.CreateBuyer(w, r)
		case "PUT":
			buyerHandler.UpdateBuyer(w, r)
		case "DELETE":
			buyerHandler.DeleteBuyer(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/sellers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			sellerHandler.GetSellers(w, r)
		case "POST":
			sellerHandler.CreateSeller(w, r)
		case "PUT":
			sellerHandler.UpdateSeller(w, r)
		case "DELETE":
			sellerHandler.DeleteSeller(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
