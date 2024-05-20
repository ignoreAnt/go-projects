package api

import (
	"encoding/json"
	"go-projects/system-stats/internal/logic"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// StatsHandler statsHandler handles the stats request
func StatsHandler(templatesDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats, err := logic.GetStats()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cpuUsageJSON, err := json.Marshal(stats.CPUUsage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmplData := struct {
			CPUUsage    string
			MemoryUsage float64
			CpuInfo     logic.CPUInfo
		}{
			CPUUsage:    string(cpuUsageJSON),
			MemoryUsage: stats.MemoryUsage,
			CpuInfo:     stats.CpuInfo,
		}
		tmplPath := filepath.Join(templatesDir, "index.html")
		tmpl := template.Must(template.ParseFiles(tmplPath))
		err = tmpl.Execute(w, tmplData)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
