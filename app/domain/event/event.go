package event

import (
	"encoding/json"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/driver/nsq"
)

type Producer nsq.Producer

type EventEmitter interface {
	Publish(topic string, message interface{}) (err error)
	DeferredPublish(topic string, at time.Time, message interface{}) (err error)
}

type eventEmitter struct {
	producer nsq.Producer
}

func NewEventEmitter(producer nsq.Producer) EventEmitter {
	return &eventEmitter{
		producer: producer,
	}
}

func (s *eventEmitter) Publish(topic string, message interface{}) (err error) {
	body, err := json.Marshal(message)

	if err != nil {
		return err
	}

	return s.producer.Publish(topic, body)
}

func (s *eventEmitter) DeferredPublish(topic string, at time.Time, message interface{}) (err error) {
	body, err := json.Marshal(message)

	if err != nil {
		return err
	}

	var delay time.Duration

	if time.Now().After(at) {
		delay = 0
	} else {
		delay = time.Until(at)
	}

	return s.producer.DeferredPublish(topic, delay, body)
}
