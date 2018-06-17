package main

import (
	"github.com/reactmed/goneurdicom/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

func main() {
	userHandler := handlers.NewUserHandler()
	router := httprouter.New()
	router.GET("/", userHandler.FindUsers)
	log.Println("Server is started")
	log.Fatal(http.ListenAndServe(":8080", router))
}