package main

import (
	"log"
	"net/http"

	setupmod "employeejwt/internal/setup.go"
)

func main() {
	setupmod.SetupDatabase()
	r := setupmod.SetupRouter()
	log.Println("Server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
