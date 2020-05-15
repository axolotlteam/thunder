package nsqueue

import (
	"fmt"
	"testing"
	"time"

	"github.com/axolotlteam/thunder/nsqueue/producer"
	"github.com/nsqio/go-nsq"
	"github.com/stretchr/testify/assert"
)

func Test_NewNSQ(t *testing.T) {
	// fmt.Println("001")
	url := "localhost:4150"

	config := producer.Configs{
		Protocol: producer.TPC,
		URL:      url,
		Retry:    10,
		Timeout:  1 * time.Second,
		LogLevel: nsq.LogLevelError,
	}

	err := NewProducer(config)
	if err != nil {
		assert.Error(t, err)
		fmt.Println(err.Error())
		return
	}

	defer producer.Queue.PublishStop()

	msg := []byte("Hello World")

	err = producer.Queue.Publish("TEST", msg)
	if err != nil {
		assert.Error(t, err)
		fmt.Println(err.Error())
		return
	}

	assert.NoError(t, err)
}
