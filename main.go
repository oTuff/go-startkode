package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oTuff/go-startkode/db"
	_ "github.com/oTuff/go-startkode/docs"
	"github.com/oTuff/go-startkode/handlers"
	"github.com/rs/cors"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func run() (*http.ServeMux, error) {
	err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/api/docs/", httpSwagger.WrapHandler)
	mux.HandleFunc("/api/todos", handlers.GetAllTodos)
	mux.HandleFunc("/api/todo/{id}", handlers.GetTodo)
	mux.HandleFunc("DELETE /api/todo/{id}", handlers.DeleteTodo)
	mux.HandleFunc("POST /api/todo", handlers.CreateTodo)

	return mux, nil
}

func main() {
	mux, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close() // Ensure the database connection is closed at the end

	handler := cors.Default().Handler(mux)

	fmt.Println("Running server on port 4000")
	log.Fatal(http.ListenAndServe(":4000", handler))
}
