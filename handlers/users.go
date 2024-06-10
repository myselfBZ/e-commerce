package handlers

import (
	errs "e-commerce/errors"
	"e-commerce/initializers"
	"e-commerce/models"
	"encoding/json"
	"net/http"
	"time"
)
     
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errs.ErrorHandle(w, http.StatusMethodNotAllowed, errs.MethodNotAllowed)
		return
	}

	var user struct {
		Name      string `json:"name"`
		LastName  string `json:"lastName"`
		BirthDate string `json:"birthDate"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errs.ErrorHandle(w, http.StatusBadRequest, errs.InvalidJson)

	}

	format := "2006-01-02"
	date, err := time.Parse(format, user.BirthDate)

	if err != nil {
		errs.ErrorHandle(w, http.StatusBadRequest, map[string]string{"Error": "Invalid date format"})
		return
	}

	if result := initializers.DB.Create(&models.User{Name: user.Name, LastName: user.LastName, BirthDate: date}); result.Error != nil {
		errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)

		return
	}
	json.NewEncoder(w).Encode(map[string]string{"Message": "Created!"})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errs.ErrorHandle(w, http.StatusMethodNotAllowed, errs.MethodNotAllowed)
		return
    
	}
	var users []models.User
	if result := initializers.DB.Find(&users); result.Error != nil {
		errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
		return
	}
	json.NewEncoder(w).Encode(users)
}
