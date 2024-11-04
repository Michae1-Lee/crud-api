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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}

	userRepo := mg.NewUserRepository(client)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	userHandler := api.NewUserHandler(userService, authService)

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("POST /register", userHandler.RegisterUser)
	authRouter.HandleFunc("POST /login", userHandler.LoginUser)

	apiRouter := http.NewServeMux()
	apiRouter.HandleFunc("GET /get/{id}", userHandler.GetUser)

	mainRouter := http.NewServeMux()
	mainRouter.Handle("POST /register", authRouter)
	mainRouter.Handle("POST /login", authRouter)
	mainRouter.Handle("GET /get/{id}", userHandler.UserIdentity(apiRouter))

	server := http.Server{
		Addr:    ":8080",
		Handler: mainRouter,
	}
	log.Println("server started")
	server.ListenAndServe()
}
