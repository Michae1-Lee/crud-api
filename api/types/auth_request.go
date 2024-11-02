package types

import (
	"encoding/json"
	"net/http"
)

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AuthToken string `json:"auth_token"`
}

func CreateRegisterRequest(r *http.Request) (*RegisterRequest, error) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func CreateLoginRequest(r *http.Request) (*LoginRequest, error) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}
