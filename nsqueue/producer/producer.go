package producer

import (
	"time"

	"github.com/bitly/go-nsq"
)

// ConnTCP -
func ConnTCP(c Configs) error {
	p, err := nsq.NewProducer(c.URL, nsq.NewConfig())
	p.SetLogger(Logger, c.LogLevel)

	if err != nil {
		return err
	}

	err = p.Ping()
	if err != nil {
		channel := make(chan bool)
		go reConnect(c.Timeout, c.Retry, p, channel)

		if !<-channel {
			p.Stop()
			return err
		}
	}

	Queue = &Tconnecter{
		URL:      c.URL,
		Producer: p,
	}

	return nil
}

// ConnHTTP -
func ConnHTTP(c Configs) {
	Queue = &Hconnecter{
		URL: c.URL,
	}
}

func reConnect(timeout time.Duration, retry int, p *nsq.Producer, c chan<- bool) {
	for x := 0; x < retry; x++ {

		if err := p.Ping(); err == nil {
			c <- true
		}

		time.Sleep(timeout)
	}
	c <- false
}
