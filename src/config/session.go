package config

import (
	"github.com/gocql/gocql"
	"sync"
)

var mu = sync.Mutex{}

func InitCluster() *gocql.Session {
	mu.Lock()
	defer mu.Unlock()
	cluster := gocql.NewCluster("localhost")
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "cassandra",
	}

	cluster.Consistency = gocql.LocalQuorum
	session, _ := cluster.CreateSession()
	return session
}
