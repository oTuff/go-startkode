package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oTuff/go-startkode/db"
)

// GetAllTodos godoc
// @Summary Get all todos
// @Description Fetches a list of all todos from the database
// @Tags todos
// @Produce application/json
// @Success 200 {array} models.Todo
// @Router /api/todos [get]
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := db.FetchAllTodos()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetTodo godoc
// @Summary Get todo
// @Description Fetches a todo based on the id from the database
// @Tags todos
// @Produce application/json
// @Param id path string true 3"Todo ID"
// @Success 200 {object} models.Todo
// @Router /api/todo/{id} [get]
func GetTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	todo, err := db.GetTodoById(id)
	if err != nil {
		log.Fatal(err)
	}

	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteTodo godoc
// @Summary Delete todo
// @Description Delete a todo based on the id from the database
// @Tags todos
// @Produce application/json
// @Param id path string true "Todo ID"
// @Success 200
// @Router /api/todo/{id} [delete]
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := db.DeleteTodoById(id)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// w.Write(res)
}
