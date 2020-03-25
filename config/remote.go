package config

import (
	"bytes"

	"github.com/axolotlteam/thunder/logger"
	"github.com/axolotlteam/thunder/st"
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
		logger.WithError(err).Panicf("connect consul failed with host: %s ", host)
		return nil, st.ErrorConnectFailed
	}

	kv := client.KV()

	pair, _, err := kv.Get(key, nil)
	if err != nil {
		logger.WithError(err).Panicf("not found the key [ %s ]exists", key)
		return nil, st.ErrorDataNotFound
	}

	switch ftype {
	case "yaml":
		v.SetConfigType("yaml")
	case "json":
		v.SetConfigType("json")
	}

	err = v.ReadConfig(bytes.NewReader(pair.Value))

	if err != nil {
		logger.WithError(err).Panicf("read consul config failed with key : %s", key)
		return nil, st.ErrorDataParseFailed
	}

	return v, nil
}
