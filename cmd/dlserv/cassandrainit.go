package main

import (
	"log"

	"github.com/dgravesa/drinklogs-service/data"
)

// creates a Cassandra client for the data backend; exits on failure.
func createCassandraClient(configName string) *data.CassandraClient {
	var client *data.CassandraClient

	// load configuration from file
	config, err := data.ReadCassandraClientConfig(configName)
	if err != nil {
		log.Fatalln(err)
	}

	// initialize client with configuration
	client, err = data.NewCassandraClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
