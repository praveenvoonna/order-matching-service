package models

type Order struct {
	ID       int     `json:"id"`
	BuyerID  int     `json:"buyer_id"`  // Foreign key referencing the Buyer
	SellerID int     `json:"seller_id"` // Foreign key referencing the Seller
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Status   string  `json:"status"` // Status of the order, e.g., "pending", "fulfilled"
}
