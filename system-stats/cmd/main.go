package main

import (
	"go-projects/system-stats/internal/api"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Get the absolute path to the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Construct absolute paths for the templates and static directories
	templatesDir := filepath.Join(currentDir, "/system-stats/templates")
	staticDir := filepath.Join(currentDir, "/system-stats/static")

	// Pass the templates directory to the StatsHandler
	http.HandleFunc("/stats", api.StatsHandler(templatesDir))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
