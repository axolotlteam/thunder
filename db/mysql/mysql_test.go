package mysql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_MySQLConnect(t *testing.T) {

	c := Config{
		Host:        "localhost:27017",
		User:        "",
		Password:    "",
		Database:    "",
		AppName:     "test",
		OpenConns:   10,
		IdleConns:   2,
		MaxLifetime: time.Second * 10,
	}

	err := Con(c)
	assert.NoError(t, err)
}
