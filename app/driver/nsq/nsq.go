package nsq

import (
	"github.com/nsqio/go-nsq"
	"time"
)

type NSQ struct {
	LookupAddress string
	Address       string
}

type Producer interface {
	DeferredPublish(topic string, delay time.Duration, body []byte) error
	Publish(topic string, message []byte) (err error)
	Stop()
}

type Consumer interface {
	AddHandler(nsq.Handler)
	ConnectToNSQD(string) error
	ConnectToNSQLookupd(string) error
	Stop()
}

type Logger interface {
	Output(calldepth int, s string) error
}

type HandlerFunc func(message *nsq.Message) error

func NewConnection(lookupaddress string, address string) *NSQ {
	return &NSQ{
		LookupAddress: lookupaddress,
		Address:       address,
	}
}

func (n *NSQ) NewProducer() (producer Producer, err error) {
	config := nsq.NewConfig()
	return nsq.NewProducer(n.Address, config)
}

func (n *NSQ) AddEventListener(topic string, channel string, handler nsq.HandlerFunc) (consumer Consumer, err error) {
	config := nsq.NewConfig()
	consumer, err = nsq.NewConsumer(topic, channel, config)

	if err != nil {
		return consumer, err
	}

	// Register handler
	consumer.AddHandler(handler)

	return consumer, nil
}
