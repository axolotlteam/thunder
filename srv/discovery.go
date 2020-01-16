package srv

import (
	"log"
	"net"
	"strconv"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/naming"
)

// Resolver implements the gRPC Resolver interface using a Consul backend.
type Resolver struct {
	c           *api.Client
	service     string
	tag         string
	passingOnly bool

	quitc    chan struct{}
	updatesc chan []*naming.Update
}

// NewResolver initializes and returns a new Resolver.
func NewResolver(client *api.Client, service, tag string) (*Resolver, error) {
	r := &Resolver{
		c:           client,
		service:     service,
		tag:         tag,
		passingOnly: true,
		quitc:       make(chan struct{}),
		updatesc:    make(chan []*naming.Update, 1),
	}

	// Retrieve instances immediately
	instances, index, err := r.getInstances(0)
	if err != nil {
		log.Printf("grpc/lb/consul: error retrieving instances from Consul: %v", err)
	}
	updates := r.makeUpdates(nil, instances)
	if len(updates) > 0 {
		r.updatesc <- updates
	}

	// Start updater
	go r.updater(instances, index)

	return r, nil
}

// by Resolver as well, see Next and Close.
func (r *Resolver) Resolve(target string) (naming.Watcher, error) {
	return r, nil
}

// Next -
func (r *Resolver) Next() ([]*naming.Update, error) {
	return <-r.updatesc, nil
}

// Close closes the watcher.
func (r *Resolver) Close() {
	select {
	case <-r.quitc:
	default:
		close(r.quitc)
		close(r.updatesc)
	}
}

func (r *Resolver) updater(instances []string, lastIndex uint64) {
	var err error
	var oldInstances = instances
	var newInstances []string

	for {
		select {
		case <-r.quitc:
			break
		default:
			newInstances, lastIndex, err = r.getInstances(lastIndex)
			if err != nil {
				log.Printf("grpc/lb/consul: error retrieving instances from Consul: %v", err)
				continue
			}
			updates := r.makeUpdates(oldInstances, newInstances)
			if len(updates) > 0 {
				r.updatesc <- updates
			}
			oldInstances = newInstances
		}
	}
}

// service from Consul.
func (r *Resolver) getInstances(lastIndex uint64) ([]string, uint64, error) {
	services, meta, err := r.c.Health().Service(r.service, r.tag, r.passingOnly, &api.QueryOptions{
		WaitIndex: lastIndex,
	})
	if err != nil {
		return nil, lastIndex, err
	}

	var instances []string
	for _, service := range services {
		s := service.Service.Address
		if len(s) == 0 {
			s = service.Node.Address
		}
		addr := net.JoinHostPort(s, strconv.Itoa(service.Service.Port))
		instances = append(instances, addr)
	}
	return instances, meta.LastIndex, nil
}

func (r *Resolver) makeUpdates(oldInstances, newInstances []string) []*naming.Update {
	oldAddr := make(map[string]struct{}, len(oldInstances))
	for _, instance := range oldInstances {
		oldAddr[instance] = struct{}{}
	}
	newAddr := make(map[string]struct{}, len(newInstances))
	for _, instance := range newInstances {
		newAddr[instance] = struct{}{}
	}

	var updates []*naming.Update
	for addr := range newAddr {
		if _, ok := oldAddr[addr]; !ok {
			updates = append(updates, &naming.Update{Op: naming.Add, Addr: addr})
		}
	}
	for addr := range oldAddr {
		if _, ok := newAddr[addr]; !ok {
			updates = append(updates, &naming.Update{Op: naming.Delete, Addr: addr})
		}
	}

	return updates
}
