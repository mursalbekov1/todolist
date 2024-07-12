package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"task2/internal/models"
	"time"
)

var (
	id        int
	title     string
	todoStore sync.Map
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id = getId(id)
	todo.ID = id
	todo.ActiveAt = time.Now().Format(time.RFC822)

	todoStore.Store(todo.ID, todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	_, ok := todoStore.Load(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	var updatedTodo models.Todo
	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTodo.ID = id
	updatedTodo.ActiveAt = time.Now().Format(time.RFC822)
	todoStore.Store(updatedTodo.ID, updatedTodo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	_, ok := todoStore.Load(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	todoStore.Delete(id)

	w.WriteHeader(http.StatusNoContent)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	value, ok := todoStore.Load(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	todo := value.(models.Todo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var todos []models.Todo

	todoStore.Range(func(key, value interface{}) bool {
		todo := value.(models.Todo)
		todos = append(todos, todo)
		return true
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getId(id int) int {
	id++
	return id
}
