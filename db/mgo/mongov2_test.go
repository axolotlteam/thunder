package mgo

import (
	"context"
	"testing"

	"github.com/axolotlteam/thunder/config"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

// Test_Con -
func Test_Con(t *testing.T) {

	config := config.Database{
		Host: "localhost:27017,localhost:27018,localhost:27019",
		//User:     "root",
		//Password: "9527",
		Database:   "admin",
		AppName:    "test",
		ReplicaSet: "rs0",
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
