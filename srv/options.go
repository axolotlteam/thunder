package srv

import (
	"log"
	"net"
	"strings"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

// Options -
type Options struct {
	id       string
	Name     string
	Server   *grpc.Server
	Host     string
	Port     string
	Listener net.Listener
	// Registry
	Registry *api.Client
	// TTL time
	TTL time.Duration
}

func newOptions(opts Options) Options {
	o := Options{}
	if opts.Name == "" {
		log.Panic("Service name is required")
	}
	o.Name = opts.Name

	if opts.Server == nil {
		o.Server = grpc.NewServer()
	} else {
		o.Server = opts.Server
	}

	if opts.Host == "" {
		o.Host = localIP()
	} else {
		o.Host = opts.Host
	}

	if opts.Port == "" {
		o.Port = "0"
	} else {
		o.Port = opts.Port
	}

	var err error
	o.Listener, err = net.Listen("tcp", o.Host+":"+o.Port)
	if err != nil {
		log.Panic(err.Error())
	}

	if opts.TTL == 0 {
		o.TTL = TTL
	} else {
		o.TTL = opts.TTL
	}

	o.Port = strings.Split(o.Listener.Addr().String(), ":")[1]

	o.Registry = newClient(opts.Registry)

	return o
}

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
