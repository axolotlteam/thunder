package srv

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
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

	port, err := strconv.Atoi(opts.Port)
	if err != nil {
		log.Panic(err.Error())
	}

	if err := opts.Registry.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      opts.id,
		Name:    opts.name,
		Address: opts.Host,
		Port:    port,
		Tags:    NewTags(opts),
		Check: &api.AgentServiceCheck{
			TTL:                            (opts.RegistryTTL + time.Second).String(),
			Timeout:                        time.Minute.String(),
			DeregisterCriticalServiceAfter: time.Minute.String(),
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
func deRegister(o *Options) error {
	o.BeforeStop()

	err := o.Registry.Agent().ServiceDeregister(o.id)

	switch ServerType {
	case GIN:
		server := &http.Server{Handler: o.Server}
		server.Close()
	case GRPC:
		o.Server.(*grpc.Server).Stop()
	}

	o.Listener.Close()

	o.AfterStop()

	return err
}
