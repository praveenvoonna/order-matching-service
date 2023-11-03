package models

type Order struct {
	ID       int
	Product  string
	Quantity int
	Price    float64
}
