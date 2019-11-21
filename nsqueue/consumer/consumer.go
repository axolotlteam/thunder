package consumer

import (
	"context"
	"fmt"

	"github.com/bitly/go-nsq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ctx      context.Context
	cancel   context.CancelFunc
	consumer *nsq.Consumer
)

// Conn -
func Conn(c Configs) (err error) {
	consumer, err = nsq.NewConsumer(
		c.Topic,
		c.Channel,
		nsq.NewConfig(),
	)
	if err != nil {
		return status.New(codes.Unavailable, "set consumer failed").Err()
	}

	consumer.SetLogger(Logger, c.LogLevel)
	ConsumerConfig = c

	return nil
}

// Start -
func Start(f nsq.HandlerFunc) error {
	states := States()
	if states.Connections != 0 {
		return nil
	}

	consumer.AddHandler(f)
	ctx, cancel = context.WithCancel(context.Background())
	if err := consumer.ConnectToNSQLookupd(ConsumerConfig.Nsqlookupd); err != nil {
		fmt.Printf(err.Error())
		return status.New(codes.Unavailable, "connect consumer failed").Err()
	}

	select {
	case <-ctx.Done():
		return nil
	}
}

// States -
func States() *nsq.ConsumerStats {
	if consumer == nil {
		return new(nsq.ConsumerStats)
	}
	return consumer.Stats()
}

// Stop -
func Stop() {
	states := States()
	if states.Connections != 0 {
		consumer.Stop()
		<-consumer.StopChan
		cancel()
	}
}
