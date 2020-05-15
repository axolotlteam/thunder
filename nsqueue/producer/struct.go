package producer

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nsqio/go-nsq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// TPC - protocol
	TPC = "tcp"
	// HTTP - protocol
	HTTP = "http"
)

var (
	Queue  NSQ
	Config Configs
	Logger = log.New(os.Stderr, "", log.Flags())
)

// Configs -
type Configs struct {
	Protocol string
	URL      string
	Retry    int
	Timeout  time.Duration
	LogLevel nsq.LogLevel
	Topic    string
}

// NSQ - interface
type NSQ interface {
	Publish(topic string, d []byte) error
	PublishStop()
}

// Tconnecter - TPC
type Tconnecter struct {
	URL      string
	Producer *nsq.Producer
}

// Hconnecter - Http
type Hconnecter struct {
	URL string
}

// Publish - TPC
func (t *Tconnecter) Publish(topic string, d []byte) error {
	if Config.Protocol != TPC {
		return nil
	}

	if err := t.Producer.Ping(); err != nil {
		if errs := ConnTCP(Config); errs != nil {
			return err
		}
	}

	if err := t.Producer.Publish(topic, d); err != nil {
		return status.New(codes.DataLoss, "nsq publish failed").Err()
	}
	return nil
}

// PublishStop -
func (t *Tconnecter) PublishStop() {
	fmt.Printf("Tconnecter PublishStop")
	t.Producer.Stop()
}

// Publish - HTTP
func (t *Hconnecter) Publish(topic string, d []byte) error {
	buf := bytes.NewReader([]byte(d))
	url := fmt.Sprintf("http://%v/pub?topic=%v", t.URL, topic)
	resp, err := http.Post(url, "application/json", buf)
	if err != nil {
		return status.New(codes.DataLoss, "nsq publish failed").Err()
	}
	defer resp.Body.Close()

	return nil
}

// PublishStop -
func (t *Hconnecter) PublishStop() {}
