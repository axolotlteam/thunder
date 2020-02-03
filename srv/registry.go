package srv

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
)

var (
	// DefaultRegistry -
	DefaultRegistry = newRegistry()
	// DefaultRegistryTTL - default
	DefaultRegistryTTL = 30 * time.Second
)

func newRegistry() *api.Client {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return client
}

func newClient(c *api.Client) *api.Client {
	if c != nil {
		return c
	}

	c, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	return c
}

func registerServer(opts *Options) {
	var h [16]byte
	rand.Read(h[:])

	opts.id = fmt.Sprintf("%s-%s", opts.name, hex.EncodeToString(h[:]))

	if err := opts.Registry.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      opts.id,
		Name:    opts.name,
		Address: opts.Host,
		Port:    opts.Port,
		Tags:    []string{opts.name, opts.id},
		Check: &api.AgentServiceCheck{
			TTL:     (opts.RegistryTTL + time.Second).String(),
			Timeout: time.Minute.String(),
		},
	}); err != nil {
		log.Panic(err.Error())
	}

	go healthCheck(*opts)
}

func healthCheck(opts Options) {
	checkid := "service:" + opts.id
	for {
		if err := opts.Registry.Agent().PassTTL(checkid, ""); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(opts.RegistryTTL)
	}
}

// DeRegister a service with consul local agent
func deRegister(c *api.Client, id string) error {
	fmt.Printf("\nDeRegister a service with consul agent: %v\n", id)
	return c.Agent().ServiceDeregister(id)
}
