package main

import (
	"context"
	"crud-api/api"
	mg "crud-api/repository/mongo"
	"crud-api/usecases/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}

	userRepo := mg.NewUserRepository(client)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	userHandler := api.NewUserHandler(userService, authService)

	router.HandleFunc("POST /register", userHandler.RegisterUser)
	router.HandleFunc("POST /login", userHandler.LoginUser)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("server started")
	server.ListenAndServe()
}
