package models

type Buyer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ProductID int    `json:"product_id"` // Foreign key referencing the Product
}
