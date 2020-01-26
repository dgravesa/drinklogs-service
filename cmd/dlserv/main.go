package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/dgravesa/drinklogs-service/auth"
	"github.com/dgravesa/drinklogs-service/controller"
	"github.com/dgravesa/drinklogs-service/data"
)

var dataBackendFlag = flag.String("data", "memory", "specify the data backend to use: [\"memory\"]")
var authBackendFlag = flag.String("authentication", "test", "authentication service to use: [\"test\"]")
var portFlag = flag.Uint("port", 33255, "port to listen on")

func main() {
	flag.Parse()
	dataBackendType := *dataBackendFlag
	authBackendType := *authBackendFlag

	var dataBackend data.DrinkLogStore
	var authBackend auth.TokenVerifier

	// initialize the data backend
	log.Printf("initializing data backend (%s)...", dataBackendType)
	switch dataBackendType {
	case "memory":
		dataBackend = data.NewInMemoryStore()
	default:
		log.Fatalf("unknown data backend type: '%s'\n", dataBackendType)
	}
	data.SetDrinkLogStore(dataBackend)
	log.Printf("data backend initialized.\n")

	// initialize the authentication backend
	log.Printf("creating authentication service (%s)...\n", authBackendType)
	switch authBackendType {
	case "test":
		authBackend = auth.NewTestTokenVerifier()
	default:
		log.Fatalf("unknown authentication backend type: '%s'\n", authBackendType)
	}
	auth.SetTokenVerifier(authBackend)
	log.Printf("authentication service initialized.\n")

	controller.InitRoutes()
	log.Println("initialized controller layer.")

	portNum := *portFlag
	log.Printf("listening on port %d...", portNum)
	http.ListenAndServe(fmt.Sprintf(":%d", portNum), nil)
}
