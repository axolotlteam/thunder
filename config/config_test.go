package config

import (
	"fmt"
	"log"
	"testing"
)

func Test_ConsulKV(t *testing.T) {
	host := "127.0.0.1:8500"
	key := "foreman"

	v, _ := ConsulKV(host, key, "yaml")

	x := struct {
		Mgo Database `json:"mgo" yaml:"mgo"`
		Log Logger   `json:"log" yaml:"log"`
	}{}

	err := v.Unmarshal(&x)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", x)
}
