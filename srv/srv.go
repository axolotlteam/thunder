package srv

import "google.golang.org/grpc"

// Service -
type Service interface {
	// The service id
	ID() string
	// The service name
	Name() string
	// Run the service
	Run() error
	// Options returns the current options
	Options() Options
	// The server
	Server() *grpc.Server
}

// NewService - creates and returns a new Service
func NewService(opts Options) Service {
	return newService(opts)
}
