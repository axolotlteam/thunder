package mgo

import (
	"context"
	"crypto/tls"
	"errors"
	"strings"
	"time"

	"github.com/axolotlteam/thunder/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var m *mongo.Client

// M -
func M() (*mongo.Client, error) {
	if m == nil {
		return nil, errors.New("no connection")
	}

	if err := m.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	return m, nil
}

// Close -
func Close() {
	m.Disconnect(context.Background())
}

// Con -
func Con(c config.Database) error {
	if err := con(c); err != nil {
		return err
	}

	if err := m.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}

	return nil
}

func con(c config.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()

	opts.ApplyURI(uri(c)).SetAppName(c.AppName)

	if c.User != "" && c.Password != "" {
		opts.SetAuth(options.Credential{
			Username:   c.User,
			Password:   c.Password,
			AuthSource: c.Database,
		})
	}
	if c.ReplicaSet != "" {
		opts.ReplicaSet = &c.ReplicaSet
	}

	if c.SSL {
		opts.SetTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})
	}

	if c.PoolSize > 0 {
		opts.SetMaxPoolSize(c.PoolSize)
	} else {
		opts.SetMaxPoolSize(10)
	}

	opts.SetMinPoolSize(5)
	opts.SetMaxConnIdleTime(time.Second * 300)
	opts.SetHeartbeatInterval(10 * time.Second)

	// Connect to MongoDB
	client, err := mongo.NewClient(opts)
	if err != nil {
		return err
	}

	if err := client.Connect(ctx); err != nil {
		return err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	m = client

	return nil
}

// mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[database][?options]]
func uri(c config.Database) string {
	// Host = "mongo1:27017,mongo2:27017"
	host := strings.Split(c.Host, ",")
	s := "mongodb://"

	for i, v := range host {
		if c.User != "" && c.Password != "" {
			s += c.User + ":" + c.Password + "@"
		}
		s += v
		if i != len(host)-1 {
			s += ","
		}
	}

	return s
}
