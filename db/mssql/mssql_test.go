package mssql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_MSSQLConnect(t *testing.T) {
	c := Config{
		Host:        "192.168.31.252",
		Port:        "27017",
		User:        "SA",
		Password:    "Axolotl!pwd123",
		Database:    "UploadEpData",
		AppName:     "ConnTest",
		OpenConns:   10,
		IdleConns:   2,
		MaxLifetime: time.Second * 10,
	}

	err := Con(c)
	assert.NoError(t, err)

}
