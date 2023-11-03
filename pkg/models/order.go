package models

type Order struct {
	ID       int
	BuyerID  int // Foreign key referencing the Buyer
	SellerID int // Foreign key referencing the Seller
	Quantity int
	Price    float64
	Status   string // Status of the order, e.g., "pending", "fulfilled"
}
