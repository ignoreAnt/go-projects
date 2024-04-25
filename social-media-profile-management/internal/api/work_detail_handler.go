package api

import (
	"encoding/json"
	"net/http"
	"social-server/internal/domain"
	"social-server/internal/service"
)

// WorkDetailsHandler is a struct that defines the handler for the work details
type WorkDetailsHandler struct {
	WorkDetailService *service.WorkDetailsService
}

// NewWorkDetailsHandler initializes a new WorkDetailsHandler
func NewWorkDetailsHandler(w *service.WorkDetailsService) *WorkDetailsHandler {
	return &WorkDetailsHandler{
		WorkDetailService: w,
	}
}

// CreateWorkDetails is a method that creates a new work detail
func (wdh WorkDetailsHandler) CreateWorkDetails(w http.ResponseWriter, r *http.Request) {

	var workDetail domain.WorkDetail

	// { WorkDetailID: 1, EmployerName: "Apple", StartYear: 2020, EndYear: 2024, JobRole: "Software Engineer", JobDescription: "Dev Ops" }
	if err := json.NewDecoder(r.Body).Decode(&workDetail); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the work detail
	if err := wdh.WorkDetailService.CreateWorkDetails(workDetail); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response to the client with the work detail id
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(map[string]int{"workDetailID": workDetail.WorkDetailID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
