package data

import "os"

import "encoding/json"

// CassandraClientConfig is used to initialize a connection to a Cassandra DB server.
type CassandraClientConfig struct {
	Hosts    []string `json:"hosts"`
	Keyspace string   `json:"keyspace"`
}

// ReadCassandraClientConfig reads the config file to get a Cassandra client configuration.
func ReadCassandraClientConfig(configName string) (CassandraClientConfig, error) {
	// open input file
	configFile, err := os.Open(configName)
	if err != nil {
		return CassandraClientConfig{}, err
	}

	// load config from file
	var config CassandraClientConfig
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)

	return config, err
}
