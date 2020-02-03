package srv

import (
	"google.golang.org/grpc"
)

// Service -
type Service interface {
	// The service id
	ID() string
	// The service name
	Name() string
	// Version
	Version() string
	// Run the service
	Run() error
	// Options returns the current options
	Options() Options
	// The server
	Server() *grpc.Server
}

// Option -
type Option func(*Options)

// NewService - creates and returns a new Service
func NewService(opts ...Option) Service {
	return newService(opts...)
}
