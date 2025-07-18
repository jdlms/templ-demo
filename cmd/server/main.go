package main

import (
	"fmt"
	"log"
	"net/http"

	"templ-demo/internal/handlers"
)

func main() {
	handlers.SetupRoutes()

	fmt.Println("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
