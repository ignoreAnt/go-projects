package api

import (
	"go-projects/social-media-profile-management/internal/repo"
	"go-projects/social-media-profile-management/internal/repo/memory"
	"go-projects/social-media-profile-management/internal/service"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	userRepository := repo.UserRepositoryFactory(true, nil)
	userService := service.NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	mux.HandleFunc("/users", userHandler.CreateUser)
	mux.HandleFunc("/users/get", userHandler.GetUserByID)

	// Picture routes setup
	pictureRepository := memory.NewPictureRepository()
	pictureService := service.NewPictureService(pictureRepository)
	pictureHandler := NewPictureHandler(pictureService)

	mux.HandleFunc("/pictures", pictureHandler.CreatePicture)
	mux.HandleFunc("/pictures/get", pictureHandler.GetPictureByID)

	// Work details routes setup
	workDetailsRepository := memory.NewWorkDetailsRepository() // Assuming factory setup
	workDetailsService := service.NewWorkDetailsService(workDetailsRepository)
	workDetailsHandler := NewWorkDetailsHandler(workDetailsService)

	mux.HandleFunc("/workDetails", workDetailsHandler.CreateWorkDetails)
}
