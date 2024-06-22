package models

import "e-commerce/initializers"



type OrderItem struct {
	ID        uint
	Product   Product `json:"product"`
	ProductID uint    `json:"productId"`
	Order     Order   `json:"order"`
	OrderID   uint    `json:"orderId"`
	Qnty      int     `json:"quantity"`
	Price     int     `json:"price"`
}

func (o *OrderItem) Create(OrItm *OrderItem) error {
    if result := initializers.DB.Model(OrItm).Create(OrItm); result.Error != nil{
        return result.Error
    }
    return nil 
}

func (o *OrderItem) Delete(id int) error {
    if result := initializers.DB.Delete(o, id); result.Error != nil{
        return result.Error
    }
    return nil 
}
