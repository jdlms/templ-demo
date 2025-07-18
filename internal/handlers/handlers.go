package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"templ-demo/internal/db"
	"templ-demo/web/templates"

	"github.com/a-h/templ"
)

var counterDB *db.CounterDB

func SetupRoutes() {
	// Initialize database
	var err error
	counterDB, err = db.NewCounterDB("counter.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Handle main page
	http.Handle("/", templ.Handler(templates.MainPage()))

	// Handle second page
	http.Handle("/second", templ.Handler(templates.SecondPage()))

	// API endpoints
	http.HandleFunc("/api/counter", handleCounter)
	http.HandleFunc("/api/counter/increment", handleCounterIncrement)

	// Serve static files (for your bundled JS and assets)
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./web/static/dist/"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/static/assets/"))))

	// Serve 3D models
	http.Handle("/models/", http.StripPrefix("/models/", http.FileServer(http.Dir("./web/static/models/"))))
}

func handleCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	count, err := counterDB.GetCounter()
	if err != nil {
		http.Error(w, "Failed to get counter", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"count": count})
}

func handleCounterIncrement(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	count, err := counterDB.IncrementCounter()
	if err != nil {
		http.Error(w, "Failed to increment counter", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"count": count})
}
