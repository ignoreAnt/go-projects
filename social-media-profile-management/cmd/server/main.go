package main

import (
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"social-server/internal/api"
	"social-server/internal/middleware"
	"social-server/pkg/log"
)

func main() {

	// Initialize the logger
	logger := log.GetLogger()

	// Set the logging level
	logger.SetLevel(log.InfoLevel)

	logger.Info("Service setup complete.")

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env.development file")
	}
	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	logger.Debug(s3Bucket, secretKey)
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
