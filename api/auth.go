package api

import (
	"crud-api/api/types"
	"crud-api/domain"
	"encoding/json"
	"net/http"
)

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateRegisterRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.userService.FindByLogin(req.Login)
	if err != nil {
		err = h.userService.CreateUser(domain.User{Login: req.Login, Password: req.Password})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

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
	if err := json.NewEncoder(w).Encode(token); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
