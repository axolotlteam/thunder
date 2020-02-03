package srv

import (
	"log"
	"net"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

// Options -
type Options struct {
	id       string
	name     string
	version  string
	Server   *grpc.Server
	Host     string
	Port     int
	Listener net.Listener
	// Registry Client
	Registry *api.Client
	// Registry TTL
	RegistryTTL time.Duration
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Server:      DefaultServer,
		Registry:    DefaultRegistry,
		Host:        DefaultHost,
		Port:        DefaultPort,
		RegistryTTL: DefaultRegistryTTL,
	}

	for _, o := range opts {
		o(&opt)
	}

	opt.Listener = newListener(&opt)

	// register server
	registerServer(&opt)

	return opt
}

// Name of the service
func Name(n string) Option {
	if n == "" {
		log.Panic("Service name is required")
	}
	return func(o *Options) {
		o.name = n
	}
}

// Version of the service
func Version(v string) Option {
	return func(o *Options) {
		o.version = v
	}
}

// Registry -
func Registry(r *api.Client) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

// RegistryTTL -
func RegistryTTL(t time.Duration) Option {
	return func(o *Options) {
		o.RegistryTTL = t
	}
}

// Server -
func Server(s *grpc.Server) Option {
	return func(o *Options) {
		o.Server = s
	}
}

// Host -
func Host(host string) Option {
	return func(o *Options) {
		o.Host = host
	}
}

// Port -
func Port(port int) Option {
	return func(o *Options) {
		o.Port = port
	}
}
