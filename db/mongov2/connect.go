package mongov2

import (
	"context"
	"time"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var m *mongo.Client

// M -
func M() (*mongo.Client, error) {
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
func Con(c Config) error {
	if err := con(c); err != nil {
		return err
	}

	if err := m.Ping(context.TODO(), readpref.Primary()); err != nil {
		spew.Dump(err.Error())
	}

	return nil
}

func con(c Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.ApplyURI(uri(c)).SetAppName(c.AppName).SetMaxPoolSize(c.PoolSize)

	if c.User != "" && c.Password != "" {
		opts.SetAuth(options.Credential{
			Username:   c.User,
			Password:   c.Password,
			AuthSource: c.Database,
		})
	}

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
func uri(c Config) string {
	s := "mongodb://"

	if c.User != "" && c.Password != "" {
		s += c.User + ":" + c.Password + "@"
	}

	s += c.Host

	return s
}
