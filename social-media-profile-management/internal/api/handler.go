package api

import (
	"encoding/json"
	"net/http"
	"social-server/internal/service"
	"strconv"
)

type ServiceHandler[T any] struct {
	profileService *service.SocialMediaProfileService[T]
}

func NewServiceHandler[T any](profileService *service.SocialMediaProfileService[T]) *ServiceHandler[T] {
	return &ServiceHandler[T]{profileService: profileService}
}

// Create creates a new item post request with the given item.
func (h *ServiceHandler[T]) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var item T

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.profileService.Create(item); err != nil {
		http.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(map[string]T{"status created": item})
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

// Get gets an item by id.
func (h *ServiceHandler[T]) Get(w http.ResponseWriter, r *http.Request) {

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
	item, err := h.profileService.Get(idInt)
	if err != nil {
		http.Error(w, "Failed to get item", http.StatusInternalServerError)
		return
	}

	// Write the item to the response
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}
