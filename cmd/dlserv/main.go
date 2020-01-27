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

var dataBackendFlag = flag.String("data", "memory", "specify the data backend to use: [\"memory\", \"cassandra\"]")
var authServiceFlag = flag.String("authentication", "test", "authentication service to use: [\"test\"]")
var portFlag = flag.Uint("port", 33255, "port to listen on")
var configNameFlag = flag.String("dbconfig", "", "config file to use when specifying a configurable data backend")

func main() {
	flag.Parse()
	dataBackendType := *dataBackendFlag
	authServiceType := *authServiceFlag
	configName := *configNameFlag

	// initialize the data backend
	log.Printf("initializing data backend (%s)...", dataBackendType)
	initializeDataBackend(dataBackendType, configName)
	log.Printf("data backend initialized.\n")

	// initialize the authentication service
	log.Printf("creating authentication service (%s)...\n", authServiceType)
	initializeAuthenticationService(authServiceType)
	log.Printf("authentication service initialized.\n")

	// initialize routes
	controller.InitRoutes()
	log.Println("initialized controller layer.")

	// listen and serve
	portNum := *portFlag
	log.Printf("listening on port %d...", portNum)
	http.ListenAndServe(fmt.Sprintf(":%d", portNum), nil)
}

func initializeDataBackend(backendType, configName string) {
	var dataBackend data.DrinkLogStore

	switch backendType {
	case "memory":
		dataBackend = data.NewInMemoryStore()
	case "cassandra":
		dataBackend = createCassandraClient(configName)
	default:
		log.Fatalf("unknown data backend type: '%s'\n", backendType)
	}

	data.SetDrinkLogStore(dataBackend)
}

func initializeAuthenticationService(serviceType string) {
	var authService auth.TokenVerifier

	switch serviceType {
	case "test":
		authService = auth.NewTestTokenVerifier()
	default:
		log.Fatalf("unknown authentication backend type: '%s'\n", serviceType)
	}

	auth.SetTokenVerifier(authService)
}
