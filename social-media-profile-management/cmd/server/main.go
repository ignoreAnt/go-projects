package main

import (
	"fmt"
	"go-projects/social-media-profile-management/internal/api"
	"go-projects/social-media-profile-management/internal/repo/memory"
	"go-projects/social-media-profile-management/internal/service"
	"go-projects/social-media-profile-management/pkg/log"
	"net/http"
)

func main() {

	// Initialize the logger
	logger := log.GetLogger()

	// Set the logging level
	logger.SetLevel(log.InfoLevel)

	// Log messages at different levels
	logger.Debug("This debug message will not appear in the log because the level is set to Info.")
	logger.Info("This is an informational message.")
	logger.Warn("This is a warning message.")
	logger.Error("This is an error message.")

	fmt.Println("Check your specified log file for output!")

	// Example setup: Determine whether to use memory based on an environment variable
	//useMemory := os.Getenv("USE_MEMORY_REPO") == "1"
	//var db *sql.DB
	//var err error
	//
	//if !useMemory {
	//	// Setup database connection
	//	db, err = sql.Open("mysql", "user:password@/dbname")
	//	if err != nil {
	//		logger.Error("Error opening database connection: ")
	//	}
	//	defer func(db *sql.DB) {
	//		err := db.Close()
	//		if err != nil {
	//			logger.Error("Error closing database connection: ")
	//		}
	//	}(db)
	//}
	//
	//// Use the factory to get the appropriate repository
	//userRepository := repo.UserRepositoryFactory(useMemory, db)

	// Setup services using the repository
	//_ = service.NewUserService(userRepository)

	logger.Info("Service setup complete.")

	userRepository := memory.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := api.NewUserHandler(userService)

	// Setup the picture service
	pictureRepository := memory.NewPictureRepository()
	pictureService := service.NewPictureService(pictureRepository)
	pictureHandler := api.NewPictureHandler(pictureService)

	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/users/get", userHandler.GetUserByID)
	// Add the picture handler
	http.HandleFunc("/pictures", pictureHandler.CreatePicture)
	http.HandleFunc("/pictures/get", pictureHandler.GetPictureByID)

	logger.Info("Starting server on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Error("Error starting server: ")
	}
}
