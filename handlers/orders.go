package handlers

import (
	errs "e-commerce/errors"
	"e-commerce/initializers"
	"e-commerce/models"
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func (h *Handler) OrdersHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateOrder(w, r)

	case http.MethodGet:
		h.GetOrders(w, r)
	
    case http.MethodDelete:
        h.DeleteOrder(w, r)
    }
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
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
    err := order.Create(&order)
    if err != nil{
        w.WriteHeader(http.StatusInternalServerError)
        return 
    }
	json.NewEncoder(w).Encode(order)
}

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
    var orders, err = h.order.GetAll()
    if err != nil{
        w.WriteHeader(http.StatusInternalServerError)
        return 
    }
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}


func (h *Handler)DeleteOrder(w http.ResponseWriter, r *http.Request){
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil{
        errs.ErrorHandle(w, http.StatusBadRequest, errs.InvalidId)
        return 
    }
    err = h.order.Delete(id)
    if err != nil{
        if err == gorm.ErrRecordNotFound{
            w.WriteHeader(http.StatusNotFound)
            return 
        }
        w.WriteHeader(http.StatusInternalServerError)
        return 
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(errs.Success)

}
