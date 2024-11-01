package main

import (
	"context"
	"crud-api/api"
	mg "crud-api/repository/mongo"
	"crud-api/usecases/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	userHandler := api.NewUserHandler(userService)

	router.HandleFunc("POST /createuser", userHandler.CreateUser)
	router.HandleFunc("GET /getuser", userHandler.GetUser)
	router.HandleFunc("POST /deleteuser", userHandler.DeleteUser)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
