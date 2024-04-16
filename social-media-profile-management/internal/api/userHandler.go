package api

import (
	"encoding/json"
	"fmt"
	"go-projects/social-media-profile-management/internal/domain"
	"go-projects/social-media-profile-management/internal/service"
	"go-projects/social-media-profile-management/pkg/log"
	"net/http"
	"strconv"
)

var logger = log.GetLogger()

// UserHandler is a struct that defines the handler for the user
type UserHandler struct {
	UserService *service.UserService
}

// NewUserHandler initializes a new UserHandler
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// CreateUser is a method that creates a new user
func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user domain.User

	// {"userID" : 1 , "userName" : "John"}
	// Decode the request body into the user struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.UserService.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]int{"userID": id})
	if err != nil {
		fmt.Println("error, ", err)
		logger.Error("Failed to encode into JSON ")
		return
	}
}

// GetUserByID is a method that gets a user by id
func (h UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}

	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}
