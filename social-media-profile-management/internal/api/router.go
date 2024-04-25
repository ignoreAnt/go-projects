package api

import (
	"net/http"
	"social-server/internal/repo"
	"social-server/internal/repo/memory"
	"social-server/internal/service"
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

func SocialMediaProfileRoutes(mux *http.ServeMux) {
	userProfileRepository := repo.GetUserProfileRepository(repo.MemoryRepo)
	userProfileService := service.NewSocialMediaProfileService(userProfileRepository)
	userProfileHandler := NewServiceHandler(userProfileService)

	mux.HandleFunc("/userProfiles", userProfileHandler.Create)
	mux.HandleFunc("/userProfiles/get", userProfileHandler.Get)
}
