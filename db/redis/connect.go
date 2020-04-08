package rds

import (
	"github.com/axolotlteam/thunder/logger"
	"github.com/go-redis/redis/v7"
)

var m *redis.Client

// M -
func M() (*redis.Client, error) {
	if _, err := m.Ping().Result(); err != nil {
		logger.WithError(err).Error("redis connect failed")
		return nil, err
	}
	return m, nil
}

// Con -
func Con(opt *redis.Options) error {
	if err := con(opt); err != nil {
		return err
	}
	return nil
}

func con(opt *redis.Options) error {
	client := redis.NewClient(opt)

	if _, err := client.Ping().Result(); err != nil {
		logger.WithError(err).Error("redis connect failed")
		return err
	}

	m = client

	return nil
}
