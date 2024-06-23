package handlers

import (
	errs "e-commerce/errors"
	"e-commerce/models"
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func (h *Handler) ProductHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetProduct(w, r)
	case http.MethodDelete:
		h.DeleteProduct(w, r)
	case http.MethodPut:
		h.UpdateProduct(w, r)
	case http.MethodPost:
		h.CreateProduct(w, r)
    
    }

}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		errs.ErrorHandle(w, http.StatusNotFound, errs.InvalidId)
		return
	}
    product, err := h.prod.Get(id)
    if err != nil{
        w.WriteHeader(http.StatusNotFound)
        errs.ErrorHandle(w, http.StatusNotFound, errs.NotFound)
        return 
    }
	json.NewEncoder(w).Encode(product)

}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var prod models.Product
    if err := json.NewDecoder(r.Body).Decode(&prod); err != nil{
        errs.ErrorHandle(w, http.StatusBadRequest, errs.InvalidJson)
        return 
    }
    err := h.prod.Create(&prod)
    if err != nil{
        errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
        return 
    }
    var msg = map[string]string{
        "message":"Created!",
    }
    json.NewEncoder(w).Encode(msg)
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	var products, err = h.prod.GetAll()
    if err != nil{
        errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
        return 
    }

	json.NewEncoder(w).Encode(products)

}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	productId, err := strconv.Atoi(id)
	if err != nil {
		errs.ErrorHandle(w, http.StatusBadRequest, errs.InvalidId)
		return
	}
	err = h.prod.Delete(productId) 
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errs.ErrorHandle(w, http.StatusNotFound, errs.NotFound)
			return
		}
		errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted successfully"})
}

func(h *Handler)  UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		var message = map[string]string{
			"message": "Invalid id",
		}
		json.NewEncoder(w).Encode(message)
		return
	}
	var newProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		errs.ErrorHandle(w, http.StatusBadRequest, errs.InvalidJson)
		return
	}
    err = h.prod.Update(id, &newProduct)	
    if err != nil{
        errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
        return
    }
	json.NewEncoder(w).Encode(map[string]string{"Message": "Updated successfully"})

}
