package config

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConsulKV(t *testing.T) {
	host := "127.0.0.1:8500"
	key := "auth"

	v, err := ConsulKV(host, key, "yaml")

	if !assert.NoError(t, err) {
		return
	}

	x := struct {
		Mgo Database `json:"mgo" yaml:"mgo"`
		Log Logger   `json:"log" yaml:"log"`
	}{}

	err = v.Unmarshal(&x)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", x)
}
