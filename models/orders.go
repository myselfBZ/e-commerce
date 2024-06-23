package models

import (
	"e-commerce/initializers"

)

type Order struct {
	ID         uint
	Price      int         `json:"price"`
	Address    Address     `json:"address" gorm:"constraint:onDelete:CASCADE;"`
	AddressID  int         `json:"addressId"`
	CustomerID int         `json:"customer"`
	OrderItems []OrderItem `gorm:"constraint:onDelete:CASCADE;" json:"orderItems"`
}

func (o *Order) Create(ordr *Order) error {
    if result := initializers.DB.Create(ordr); result.Error != nil{
        return result.Error
    }
    return nil 
}


func (o *Order) Update(newOrd *Order, id int) error {
    var oldOrdr *Order
    if result := initializers.DB.First(oldOrdr, id); result.Error != nil{
        return result.Error
    }
    updates := initializers.DB.Model(oldOrdr).Updates(newOrd)
    if updates.Error != nil{
        return updates.Error 
    }

    return nil 
}


func (o *Order) Delete(id int) error {
    if result := initializers.DB.Delete(o, id); result.Error != nil{
        return result.Error 
    }
    return nil 
} 

func (o *Order) Get(id int) (*Order, error) {
    var order *Order 
    if result := initializers.DB.First(order, id); result.Error != nil{
        return nil, result.Error
    }
    return order, nil 
}


func (o *Order) GetAll() ([]Order, error) {
    var Orders []Order 
    if result := initializers.DB.Preload("products, users, order_items").Model(o).Find(&Orders); result.Error != nil{
        return nil, result.Error 
    }
    return Orders, nil  
}



