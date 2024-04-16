package main

import (
	"fmt"
	"go-projects/social-media-profile-management/pkg/log"
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
}
