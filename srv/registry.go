package srv

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/consul/api"
)

const (
	// TTL - default
	TTL = 30 * time.Second
)

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

	opts.id = fmt.Sprintf("%s-%s", opts.Name, hex.EncodeToString(h[:]))

	port, err := strconv.Atoi(opts.Port)
	if err != nil {
		log.Panic(err.Error())
	}
	spew.Dump(opts.TTL)
	if err := opts.Registry.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      opts.id,
		Name:    opts.Name,
		Address: opts.Host,
		Port:    port,
		Tags:    []string{opts.Name, opts.id},
		Check: &api.AgentServiceCheck{
			TTL:     (opts.TTL + time.Second).String(),
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
		time.Sleep(opts.TTL)
	}
}

// DeRegister a service with consul local agent
func deRegister(c *api.Client, id string) error {
	fmt.Printf("\nDeRegister a service with consul agent: %v\n", id)
	return c.Agent().ServiceDeregister(id)
}
