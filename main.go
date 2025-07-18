package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	// Handle main page
	http.Handle("/", templ.Handler(mainPage()))

	// Handle second page
	http.Handle("/second", templ.Handler(secondPage()))

	// Serve static files (for your bundled JS)
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist/"))))

	// Serve 3D models
	http.Handle("/models/", http.StripPrefix("/models/", http.FileServer(http.Dir("./models/"))))

	fmt.Println("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
