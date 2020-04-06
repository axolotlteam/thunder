package srv

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

// ServerType -
var ServerType int

// server type list
const (
	GRPC = iota
	GIN
)

// Options -
type Options struct {
	id       string
	name     string
	version  string
	Server   http.Handler
	Host     string
	Port     string
	Listener net.Listener
	// Registry Client
	Registry *api.Client
	// Registry TTL
	RegistryTTL time.Duration

	// Before and After funcs
	BeforeStart func()
	BeforeStop  func()
	AfterStart  func()
	AfterStop   func()
}

func newOptions(opts ...Option) Options {
	opt := Options{
		Server:      DefaultServer,
		Registry:    DefaultRegistry,
		Host:        DefaultHost,
		Port:        DefaultPort,
		RegistryTTL: DefaultRegistryTTL,
		BeforeStart: DefaultHook,
		BeforeStop:  DefaultHook,
		AfterStart:  DefaultHook,
		AfterStop:   DefaultHook,
	}

	for _, o := range opts {
		o(&opt)
	}

	opt.Listener = newListener(&opt)

	return opt
}

// SetName of the service
func SetName(n string) Option {
	if n == "" {
		log.Panic("Service name is required")
	}
	return func(o *Options) {
		o.name = n
	}
}

// SetVersion of the service
func SetVersion(v string) Option {
	return func(o *Options) {
		o.version = v
	}
}

// SetRegistry -
func SetRegistry(r *api.Client) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

// SetRegistryTTL -
func SetRegistryTTL(t time.Duration) Option {
	return func(o *Options) {
		o.RegistryTTL = t
	}
}

// SetServer -
func SetServer(s http.Handler) Option {
	return func(o *Options) {
		switch s.(type) {
		case *gin.Engine:
			ServerType = GIN
		default:
			ServerType = GRPC
		}
		o.Server = s
	}
}

// SetHost -
func SetHost(host string) Option {
	return func(o *Options) {
		o.Host = host
	}
}

// SetPort -
func SetPort(port string) Option {
	return func(o *Options) {
		o.Port = port
	}
}

// Before and Afters

// BeforeStart -
func BeforeStart(fn func()) Option {
	return func(o *Options) {
		o.BeforeStart = fn
	}
}

// BeforeStop -
func BeforeStop(fn func()) Option {
	return func(o *Options) {
		o.BeforeStop = fn
	}
}

// AfterStart -
func AfterStart(fn func()) Option {
	return func(o *Options) {
		o.AfterStart = fn
	}
}

// AfterStop -
func AfterStop(fn func()) Option {
	return func(o *Options) {
		o.AfterStop = fn
	}
}
