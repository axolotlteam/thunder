package consumer

import (
	"log"
	"os"

	"github.com/nsqio/go-nsq"
)

var (
	Logger         = log.New(os.Stderr, "", log.Flags())
	ConsumerConfig Configs
)

// Configs -
type Configs struct {
	LogLevel   nsq.LogLevel
	Topic      string
	Channel    string
	Nsqlookupd string
}
