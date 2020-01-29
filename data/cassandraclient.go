package data

import (
	"time"

	"github.com/dgravesa/drinklogs-service/model"

	"github.com/gocql/gocql"
)

// CassandraClient is a data store implementation using a Cassandra DB backend.
type CassandraClient struct {
	session *gocql.Session
}

// NewCassandraClient creates a new Cassandra DB client.
func NewCassandraClient(config CassandraClientConfig) (*CassandraClient, error) {
	var client CassandraClient
	var err error

	// create session to keyspace
	cluster := gocql.NewCluster(config.Hosts...)
	cluster.Keyspace = config.Keyspace
	cluster.Consistency = gocql.Quorum
	client.session, err = cluster.CreateSession()

	return &client, err
}

// Insert creates a new drink log for a user.
func (c *CassandraClient) Insert(uid uint64, log model.DrinkLog) error {
	return c.session.Query(
		`INSERT INTO drinklogs (uid, time, amount) VALUES (?, ?, ?)`,
		uid, gocql.UUIDFromTime(log.Time), log.Amount).Exec()
}

// InRange returns drink logs within a specified time range for a user.
func (c *CassandraClient) InRange(uid uint64, ti, tf time.Time) ([]model.DrinkLog, error) {
	var reslogs []model.DrinkLog

	// execute query
	iter := c.session.Query(
		`SELECT toTimestamp(time), amount FROM drinklogs
		WHERE uid = ? AND time >= ? AND time <= ?`,
		uid, gocql.MinTimeUUID(ti), gocql.MaxTimeUUID(tf)).Consistency(gocql.One).Iter()

	// get logs from query result
	var log model.DrinkLog
	for iter.Scan(&log.Time, &log.Amount) {
		reslogs = append(reslogs, log)
	}
	err := iter.Close()

	return reslogs, err
}
