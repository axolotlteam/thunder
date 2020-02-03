package srv

import (
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

type service struct {
	opts Options
}

// newService -
func newService(opts ...Option) Service {
	options := newOptions(opts...)

	return &service{
		opts: options,
	}
}

func (s *service) ID() string {
	return s.opts.id
}

func (s *service) Version() string {
	return s.opts.version
}

func (s *service) Name() string {
	return s.opts.name
}

func (s *service) Options() Options {
	return s.opts
}

func (s *service) Server() *grpc.Server {
	return s.opts.Server
}

func (s *service) Run() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	go func(o *Options) {
		<-c
		deRegister(o.Registry, o.id)
		os.Exit(1)
	}(&s.opts)

	return s.Start()
}

func (s *service) Start() error {
	if err := s.opts.Server.Serve(s.opts.Listener); err != nil {
		return err
	}
	return nil
}
