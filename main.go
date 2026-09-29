package main

import (
	"log"
	"net/http"
	"os"

	setupmod "employeejwt/internal/setup.go"
)

func main() {
	setupmod.SetupDatabase()
	r := setupmod.SetupRouter()
	log.Println(os.Getenv("SERVER_PROTOCOL"), "server started on", os.Getenv("SERVER_HOST")+":"+os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), r))
}
