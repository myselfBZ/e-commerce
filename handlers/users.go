package handlers

import (
	errs "e-commerce/errors"
	"e-commerce/models"
	"encoding/json"
	"net/http"
	"time"
)

func(h *Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetUsers(w, r)
	case http.MethodPost:
		h.CreateUser(w, r)
	}
}

func (h *Handler)CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err := h.user.Create(&models.User{Name: user.Name, LastName: user.LastName, BirthDate: date}); err != nil {
		errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)

		return
	}
	json.NewEncoder(w).Encode(map[string]string{"Message": "Created!"})
}

func (h *Handler)GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errs.ErrorHandle(w, http.StatusMethodNotAllowed, errs.MethodNotAllowed)
		return

	}
	var users, err = h.user.GetUsers()
	if err != nil {
		errs.ErrorHandle(w, http.StatusInternalServerError, errs.InternalServer)
		return
	}
	json.NewEncoder(w).Encode(users)
}
