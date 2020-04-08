package rds

import (
	"testing"

	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/assert"
)

func Test_RedisConnect(t *testing.T) {

	c := &redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10,
		MaxRetries:   3,
		MinIdleConns: 5,
	}

	err := Con(c)
	assert.NoError(t, err)

	client, _ := M()
	res := client.Set("123", "456", 0)
	assert.Equal(t, "OK", res.Val())

}
