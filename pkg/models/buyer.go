package models

type Buyer struct {
	ID        int
	Name      string
	Email     string
	Budget    float64
	Quantity  int
	ProductID int // Foreign key referencing the Product
}
