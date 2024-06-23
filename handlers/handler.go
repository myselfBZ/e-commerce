package handlers

import "e-commerce/models"


type Handler struct{
    prod *models.Product
    user *models.User 
    order *models.Order
    orItm *models.OrderItem
    address *models.Address
}

func NewHandler(prod *models.Product, user *models.User, order *models.Order, orItm *models.OrderItem) *Handler{
    return &Handler{
        prod: prod,
        user: user,
        order: order,
        orItm: orItm,
    }
}
