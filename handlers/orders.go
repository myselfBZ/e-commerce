package handlers

import (
	errs "e-commerce/errors"
	"e-commerce/initializers"
	"e-commerce/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errs.ErrorHandle(w, http.StatusMethodNotAllowed, errs.MethodNotAllowed)
		return
	}

	var order models.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		errs.ErrorHandle(w, http.StatusBadRequest, errs.InvalidJson)
		return
	}
	for i := range order.OrderItems {
		var product models.Product
		item := &order.OrderItems[i]
		if result := initializers.DB.Where("id = ?", item.ProductID).First(&product); result.Error != nil {
			errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
			return
		}
		fmt.Println(product)
		if product.ID == 0 {
			errs.ErrorHandle(w, http.StatusNotFound, errs.NotFound)

			return
		}
		if product.CountInStock == 0 {
			errs.ErrorHandle(w, http.StatusNotFound, errs.InternalServer)
			return
		}
		countInStock := &product.CountInStock
		item.Price = item.Qnty * product.Price
		fmt.Println(item.Price)
		*countInStock -= item.Qnty
		fmt.Println(product.CountInStock)
		if result := initializers.DB.Save(&product); result.Error != nil {
			errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
			return
		}
	}
	fmt.Println(order.OrderItems)

	if result := initializers.DB.Create(&order); result.Error != nil {
		errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
		return
	}
	json.NewEncoder(w).Encode(order)
}
