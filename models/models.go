package models

import "time"

type Product struct {
	ID           uint   `gorm:"primaryKey" json:"productId"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	CountInStock int    `json:"countInStck"`
	Description  string `json:"description"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	BirthDate time.Time `json:"birthDate"`
}

type Order struct {
	ID         uint
	Price      int         `json:"price"`
	Address    Address     `json:"address"`
	AddressID  int         `json:"addressId"`
	CustomerID int         `json:"customer"`
	OrderItems []OrderItem `json:"orderItems"`
}

type Address struct {
	ID          int
	City        string `gorm:"type:varchar(50)" json:"city"`
	Country     string `gorm:"type:varchar(50)" json:"country"`
	AddressLine string `gorm:"type:varchar(50)" json:"addressLine"`
}

type OrderItem struct {
	ID        uint
	Product   Product `json:"product"`
	ProductID uint    `json:"productId"`
	Order     Order   `json:"order"`
	OrderID   uint    `json:"orderId"`
	Qnty      int     `json:"quantity"`
	Price     int     `json:"price"`
}
