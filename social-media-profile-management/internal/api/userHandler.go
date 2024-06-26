package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-server/internal/domain"
	"social-server/internal/service"
	"social-server/pkg/log"
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

	if r.Method != "POST" {
		http.Error(w, "Invalid HTTP method Type", http.StatusBadRequest)
		return
	}

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
	// Get the user id from the query parameter
	idStr := r.URL.Query().Get("userID")
	// Convert the id to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user Id", http.StatusBadRequest)
		return
	}

	// Get the user by id
	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}
