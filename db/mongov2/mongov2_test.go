package mongov2

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

// Test_Con -
func Test_Con(t *testing.T) {

	config := Config{
		Host:     "localhost:27017",
		User:     "",
		Password: "",
		Database: "testDatabase",
		AppName:  "test",
	}

	err := Con(config)
	assert.NoError(t, err)

	client, err := M()
	assert.NoError(t, err)

	c := client.Database("test").Collection("testConfig")

	cur, err := c.Find(context.Background(), bson.M{})
	assert.NoError(t, err)

	for cur.Next(context.Background()) {
		res := bson.M{}
		cur.Decode(&res)
		spew.Dump(res)
	}
}
