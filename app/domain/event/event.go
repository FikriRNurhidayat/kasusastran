package event

import (
	"encoding/json"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/driver/nsq"
	"github.com/google/uuid"
)

type Producer nsq.Producer

type Message struct {
	ID        uuid.UUID   `json:"id"`
	Kind      string      `json:"kind"`
	CreatedAt time.Time   `json:"created_at"`
	Actor     *Actor      `json:"actor"`
	Payload   interface{} `json:"payload"`
}

type Actor struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
}

type Publisher interface {
	Publish(topic string, message interface{}) (err error)
	DeferredPublish(topic string, at time.Time, message interface{}) (err error)
}

type EventEmitter struct {
	producer nsq.Producer
}

func (s *EventEmitter) Publish(topic string, message interface{}) (err error) {
	body, err := json.Marshal(message)

	if err != nil {
		return err
	}

	return s.producer.Publish(topic, body)
}

func (s *EventEmitter) DeferredPublish(topic string, at time.Time, message interface{}) (err error) {
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
