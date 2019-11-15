package nsqueue

import (
	"github.com/axolotlteam/thunder/nsqueue/consumer"
	"github.com/axolotlteam/thunder/nsqueue/producer"
)

// NewProducer -
func NewProducer(c producer.Configs) error {
	switch c.Protocol {
	case producer.TPC:
		producer.Config = c
		if err := producer.ConnTCP(c); err != nil {
			return err
		}
	case producer.HTTP:
		producer.ConnHTTP(c)
	}
	return nil
}

// NewConsumer -
func NewConsumer(c consumer.Configs) error {
	if err := consumer.Conn(c); err != nil {
		return err
	}
	return nil
}
