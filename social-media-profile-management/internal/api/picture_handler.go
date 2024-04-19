package api

import (
	"encoding/json"
	"go-projects/social-media-profile-management/internal/domain"
	"go-projects/social-media-profile-management/internal/service"
	"net/http"
	"strconv"
)

// PictureHandler is a struct that defines the handler for the picture
type PictureHandler struct {
	PictureService *service.PictureService
}

// NewPictureHandler initializes a new PictureHandler
func NewPictureHandler(pictureService *service.PictureService) *PictureHandler {
	return &PictureHandler{
		PictureService: pictureService,
	}
}

// CreatePicture is a method that creates a new picture
func (h PictureHandler) CreatePicture(w http.ResponseWriter, r *http.Request) {

	var picture domain.Picture

	// {"pictureID" : 1 , "pictureURL" : "http://www.google.com", "size" : 100, "pictureType" : "cover"}
	// Decode the request body into the picture struct
	if err := json.NewDecoder(r.Body).Decode(&picture); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.PictureService.CreatePicture(picture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"pictureURL": picture.PictureURL})
	if err != nil {
		logger.Error("Failed to encode into JSON ")
		return
	}
}

// GetPictureByID is a method that gets a picture by id
func (h PictureHandler) GetPictureByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid picture Id", http.StatusBadRequest)
		return
	}

	picture, err := h.PictureService.GetPictureByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(picture)
	if err != nil {
		return
	}

}
