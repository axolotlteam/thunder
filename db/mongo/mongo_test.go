package mongo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MongoConnect(t *testing.T) {

	c := Config{
		Host:     "localhost:27017",
		User:     "",
		Password: "",
		Database: "",
		AppName:  "test",
	}

	err := Con(c)
	assert.NoError(t, err)

}
