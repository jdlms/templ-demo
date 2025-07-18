package handlers

import (
	"net/http"

	"templ-demo/web/templates"

	"github.com/a-h/templ"
)

func SetupRoutes() {
	// Handle main page
	http.Handle("/", templ.Handler(templates.MainPage()))

	// Handle second page
	http.Handle("/second", templ.Handler(templates.SecondPage()))

	// Serve static files (for your bundled JS and assets)
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./web/static/dist/"))))

	// Serve 3D models
	http.Handle("/models/", http.StripPrefix("/models/", http.FileServer(http.Dir("./web/static/models/"))))
}
