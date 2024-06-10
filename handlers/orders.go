package handlers

import (
	errs "e-commerce/errors"
	"e-commerce/initializers"
	"e-commerce/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func OrdersHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateOrder(w, r)

	case http.MethodGet:
		GetOrders(w, r)
	}
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		errs.ErrorHandle(w, http.StatusBadRequest, errs.InvalidJson)
		return
	}
	for i := range order.OrderItems {
		var product models.Product
		item := &order.OrderItems[i]
		if result := initializers.DB.Where("id = ?", item.ProductID).First(&product); result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				errs.ErrorHandle(w, http.StatusNotFound, errs.NotFound)
				return
			}
			errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
			return
		}

		if product.CountInStock == 0 {
			errs.ErrorHandle(w, http.StatusNotFound, errs.NotFound)
			return
		}
		item.Price = item.Qnty * product.Price
		product.CountInStock -= item.Qnty
		if result := initializers.DB.Save(&product); result.Error != nil {
			errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
			return
		}
	}

	if result := initializers.DB.Create(&order); result.Error != nil {
		errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
		return
	}
	json.NewEncoder(w).Encode(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	if result := initializers.DB.Find(&orders); result.Error != nil {
		errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}
