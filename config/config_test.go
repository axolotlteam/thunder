package config

import (
	"testing"
)

func Test_ConsulKV(t *testing.T) {
	host := "127.0.0.1:8500"
	key := "foreman"
	key2 := "foreman-yaml"

	ConsulKV(host, key, "json")
	ConsulKV(host, key2, "yaml")

}
