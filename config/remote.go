package config

import (
	"bytes"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

// ConsulKV -
func ConsulKV(host, key string, ftype string) (*viper.Viper, error) {
	v := viper.New()

	c := &api.Config{
		Address: host,
		Scheme:  "http",
	}

	client, err := api.NewClient(c)
	if err != nil {
		return nil, err
	}

	kv := client.KV()

	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return nil, err
	}

	switch ftype {
	case "yaml":
		v.SetConfigType("yaml")
	case "json":
		v.SetConfigType("json")
	}

	err = v.ReadConfig(bytes.NewReader(pair.Value))

	if err != nil {
		return nil, err
	}

	return v, nil
}
