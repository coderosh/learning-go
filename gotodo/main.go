package main

import (
	"log"
	"net/http"

	"gotodo/database"
	"gotodo/handlers"
	"gotodo/middlewares"
)

func main() {
	database.InitGlobalDB()

	httpHandler := getHttpHandler()
	handler := middlewares.Logger(httpHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Println("Server Listening on Port 8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getHttpHandler() *http.ServeMux {
	pageHandler := handlers.PageHandler{}
	todoHandler := handlers.TodoHandler{}

	mux := http.NewServeMux()

	mux.HandleFunc("/", pageHandler.GetHomePage())
	mux.HandleFunc("GET /todos", todoHandler.GetTodos())
	mux.HandleFunc("POST /todos", todoHandler.CreateTodo())
	mux.HandleFunc("POST /todos/{id}", todoHandler.DeleteTodo())

	return mux
}
