package data

// CassandraClientConfig is used to initialize a connection to a Cassandra DB server.
type CassandraClientConfig struct {
}

// ReadCassandraClientConfig reads the config file to get a Cassandra client configuration.
func ReadCassandraClientConfig(configName string) (CassandraClientConfig, error) {
	// TODO implement
	return CassandraClientConfig{}, nil
}
