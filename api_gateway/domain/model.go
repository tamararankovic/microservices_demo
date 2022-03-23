package domain

import "time"

type Product struct {
	Name  string
	Color string
}

type OrderItem struct {
	Product  Product
	Quantity uint16
}

type OrderDetails struct {
	Id              string
	CreatedAt       time.Time
	Status          string
	ShippingAddress string
	ShippingStatus  string
	OrderItems      []OrderItem
}
