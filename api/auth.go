package api

import (
	"crud-api/api/types"
	"crud-api/domain"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateRegisterRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = h.userService.FindByLogin(req.Login)
	if err == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	newUuid := uuid.New()
	err = h.userService.CreateUser(domain.User{Id: newUuid.String(), Login: req.Login, Password: req.Password})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateLoginRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.authService.GenerateToken(req.Login, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	loginResponse := types.LoginResponse{AuthToken: token}
	if err := json.NewEncoder(w).Encode(loginResponse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
