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

func ProductHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetProduct(w, r)
	case http.MethodDelete:
		DeleteProduct(w, r)
	case http.MethodPut:
		UpdateProduct(w, r)
	case http.MethodPost:
		CreateProduct(w, r)
    
    }

}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		errs.ErrorHandle(w, http.StatusNotFound, errs.InvalidId)
		return
	}
	var product models.Product
	if result := initializers.DB.First(&product, id); result.Error != nil {
		errs.ErrorHandle(w, 404, errs.NotFound)
		return
	}
	json.NewEncoder(w).Encode(product)

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		errs.ErrorHandle(w, http.StatusBadRequest, errs.InvalidJson)
		return
	}
	result := initializers.DB.Create(&product)
	if result.Error != nil {
		errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
		return
	}
	var message = map[string]string{
		"Message": "Created",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	initializers.DB.Find(&products)

	json.NewEncoder(w).Encode(products)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	productId, err := strconv.Atoi(id)
	if err != nil {
		errs.ErrorHandle(w, http.StatusBadRequest, errs.InvalidId)
		return
	}
	result := initializers.DB.Delete(&models.Product{}, productId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			errs.ErrorHandle(w, http.StatusNotFound, errs.NotFound)
			return
		}
		errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted successfully"})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
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
	var oldProduct models.Product
	result := initializers.DB.First(&oldProduct, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			errs.ErrorHandle(w, http.StatusNotFound, errs.NotFound)
			return
		}
		errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
		return
	}
	if result := initializers.DB.Model(&oldProduct).Updates(models.Product{
		Name:         newProduct.Name,
		Description:  newProduct.Description,
		Price:        newProduct.Price,
		CountInStock: newProduct.CountInStock,
	}); result.Error != nil {
		errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
	}
	json.NewEncoder(w).Encode(map[string]string{"Message": "Updated successfully"})

}
