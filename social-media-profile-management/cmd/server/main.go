package main

import (
	"fmt"
	"go-projects/social-media-profile-management/internal/api"
	"go-projects/social-media-profile-management/internal/middleware"
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
	// Create a new ServeMux
	mux := http.NewServeMux()
	// Register the routes
	api.RegisterRoutes(mux)

	// Wrap the ServeMux with the logger middleware
	loggedRouter := middleware.LoggerMiddleWare(mux)

	logger.Info("Starting server on port 8080")

	// Start the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedRouter,
	}
	if err := server.ListenAndServe(); err != nil {
		logger.Error("Error starting server: ")
	}
}
