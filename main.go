package main

import (
	"log"
	"net"
	"net/http"
	"os"

	setupmod "employeejwt/internal/setup.go"
)

func main() {
	setupmod.SetupDatabase()
	r := setupmod.SetupRouter()
	serverProtocol := os.Getenv("SERVER_PROTOCOL")

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	listenAddress := net.JoinHostPort(serverHost, serverPort)
	displayHost := serverHost

	serverURL := serverProtocol + "://" + net.JoinHostPort(displayHost, serverPort)
	os.Setenv("SERVER_PROTOCOL", serverProtocol)
	log.Println("Server Protocol:", os.Getenv("SERVER_PROTOCOL"))
	log.Println("Server Host:", serverHost)
	log.Println("Server Port:", serverPort)
	log.Println("Server URL:", serverURL)
	log.Fatal(http.ListenAndServe(listenAddress, r))
}
