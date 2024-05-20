package api

import (
	"context"
	"encoding/json"
	"net/http"
	"social-server/internal/service"
	"strconv"
)

type ServiceHandler[T any] struct {
	genericService *service.GenericService[T]
}

func NewServiceHandler[T any](ser *service.GenericService[T]) *ServiceHandler[T] {
	return &ServiceHandler[T]{genericService: ser}
}

// CreateHandler Create creates a new item post request with the given item.
func (h *ServiceHandler[T]) CreateHandler(w http.ResponseWriter, r *http.Request) {

	var entity T

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	createdEntity, err := h.genericService.Create(context.Background(), entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(createdEntity)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

// GetHandler Get gets an item by id.
func (h *ServiceHandler[T]) GetHandler(w http.ResponseWriter, r *http.Request) {

	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get the id parameter from the URL
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// Convert the id to an integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Failed to convert id to int", http.StatusBadRequest)
		return
	}

	// Get the item from backend
	item, err := h.genericService.GetById(context.Background(), idInt)
	if err != nil {
		http.Error(w, "Failed to get item", http.StatusInternalServerError)
		return
	}

	// Write the item to the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

// UpdateHandler Update updates an item by id.
func (h *ServiceHandler[T]) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	var entity T

	// Check if the request method is PUT
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	updatedEntity, err := h.genericService.Update(context.Background(), entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(updatedEntity)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// DeleteHandler Delete deletes an item by id.
func (h *ServiceHandler[T]) DeleteHandler(w http.ResponseWriter, r *http.Request) {

	// Check if the request method is DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get the id parameter from the URL
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// Convert the id to an integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Failed to convert id to int", http.StatusBadRequest)
		return
	}

	// Delete the item from backend
	err = h.genericService.Delete(context.Background(), idInt)
	if err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
