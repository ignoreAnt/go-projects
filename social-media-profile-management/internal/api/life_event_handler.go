package api

import (
	"encoding/json"
	"go-projects/social-media-profile-management/internal/domain"
	"go-projects/social-media-profile-management/internal/service"
	"net/http"
	"strconv"
)

type LifeEventHandler struct {
	LifeEventService *service.LifeEventService
}

// NewLifeEventHandler is a function that returns a new LifeEventHandler
func NewLifeEventHandler(service *service.LifeEventService) *LifeEventHandler {
	return &LifeEventHandler{LifeEventService: service}
}

// Create is a method that creates a new life event
func (h *LifeEventHandler) CreateLifeEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var lifeEvent domain.LifeEvent

	// { LifeEventID: 1, EventName: "Graduation", EventType: "Education", EventDate: "2020-06-01" }
	err := json.NewDecoder(r.Body).Decode(&lifeEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.LifeEventService.Create(lifeEvent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err2 := json.NewEncoder(w).Encode(map[string]int{"lifeEventID": lifeEvent.LifeEventID})
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

// Update is a method that updates a life event
func (h *LifeEventHandler) UpdateLifeEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var lifeEvent domain.LifeEvent

	err := json.NewDecoder(r.Body).Decode(&lifeEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.LifeEventService.Update(lifeEvent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err2 := json.NewEncoder(w).Encode(map[string]int{"lifeEventID": lifeEvent.LifeEventID})
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Delete is a method that deletes a life event
func (h *LifeEventHandler) DeleteLifeEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get lifeEventid from the URL query parameter
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.LifeEventService.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
