package handlers

import (
	errs "e-commerce/errors"
	"e-commerce/initializers"
	"e-commerce/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errs.ErrorHandle(w, http.StatusMethodNotAllowed, errs.MethodNotAllowed)
		return
	}
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
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		var message = map[string]string{
			"Error": "Method not allowerd",
		}
		json.NewEncoder(w).Encode(message)
		return
	}
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var message = map[string]string{
			"message": "Invalid data",
		}
		json.NewEncoder(w).Encode(message)
		r.Body.Close()
		return
	}
	defer r.Body.Close()
	log.Println(product)
	result := initializers.DB.Create(&product)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(result.Error)
		return
	}
	var message = map[string]string{
		"Message": "Created",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		var message = map[string]string{
			"Message": "Method not allowed",
		}
		json.NewEncoder(w).Encode(message)
		return
	}
	var products []models.Product
	initializers.DB.Find(&products)

	json.NewEncoder(w).Encode(products)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	id := r.PathValue("id")

	productId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid id")
		return
	}
	result := initializers.DB.Delete(&models.Product{}, productId)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal server error")
		return
	}
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted successfully")

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		var message = map[string]string{
			"Error": "Method Not allowed",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(message)
		return
	}
	fmt.Println("Request is recieved")
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
		var message = map[string]string{
			"error": "Invalid data",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		r.Body.Close()
		return
	}
	defer r.Body.Close()
	var oldProduct models.Product
	result := initializers.DB.First(&oldProduct, id)

	if result.Error != nil {
		var message = map[string]string{
			"Error": "Not found",
		}
		errs.ErrorHandle(w, 400, message)
		return
	}
	if result := initializers.DB.Model(&oldProduct).Updates(models.Product{
		Name:         newProduct.Name,
		Description:  newProduct.Description,
		Price:        newProduct.Price,
		CountInStock: newProduct.CountInStock,
	}); result.Error != nil {
		fmt.Println("Error in the database ")
	}
	if result.RowsAffected == 0 {
		var message = map[string]string{
			"Error": "Not found",
		}
		errs.ErrorHandle(w, 404, message)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"Message": "Updated successfully"})

}
