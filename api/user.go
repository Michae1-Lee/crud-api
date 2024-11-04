package api

import (
	"crud-api/domain"
	"crud-api/usecases"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService usecases.UserServiceInterface
	authService usecases.AuthServiceInterface
}

func NewUserHandler(userService usecases.UserServiceInterface, authService usecases.AuthServiceInterface) *UserHandler {
	return &UserHandler{userService, authService}
}
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "JSON Decode Error: "+err.Error(), http.StatusBadRequest)
		fmt.Println("Decode Error:", err)
		return
	}

	if err := h.userService.CreateUser(user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := h.userService.GetUser(id)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var id int
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.userService.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
