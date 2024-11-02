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

func run() (http.Handler, error) {
	err := db.ConnectDB()
	if err != nil {
		return nil, err
	}
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("GET /api/docs/", httpSwagger.WrapHandler)
	mux.HandleFunc("GET /api/todos", handlers.GetAllTodos)
	mux.HandleFunc("GET /api/todo/{id}", handlers.GetTodo)
	mux.HandleFunc("DELETE /api/todo/{id}", handlers.DeleteTodo)
	mux.HandleFunc("POST /api/todo", handlers.CreateTodo)

	//CORS stuff
	handler := cors.Default().Handler(mux)

	return handler, err
}

func main() {
	mux, err := run()
	if err != nil {
		log.Fatal(err)
	}
	// Ensure the database connection is closed at the end
	defer db.DB.Close()

	fmt.Println("Running server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
