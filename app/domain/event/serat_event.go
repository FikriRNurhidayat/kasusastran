package event

import (
	"github.com/fikrirnurhidayat/kasusastran/app/driver/nsq"
)

type SeratEventEmitter interface {
	EmitCreatedEvent(message *Message) error
	EmitRetrievedEvent(message *Message) error
	EmitListedEvent(message *Message) error
	EmitUpdatedEvent(message *Message) error
	EmitDeletedEvent(message *Message) error
}

type seratEventEmitter struct {
	eventEmitter EventEmitter
}

func (e *seratEventEmitter) EmitCreatedEvent(message *Message) error {
	return e.eventEmitter.Publish(SERAT_CREATED_TOPIC, message)
}

func (e *seratEventEmitter) EmitDeletedEvent(message *Message) error {
	return e.eventEmitter.Publish(SERAT_DELETED_TOPIC, message)
}

func (e *seratEventEmitter) EmitListedEvent(message *Message) error {
	return e.eventEmitter.Publish(SERAT_LISTED_TOPIC, message)
}

func (e *seratEventEmitter) EmitRetrievedEvent(message *Message) error {
	return e.eventEmitter.Publish(SERAT_RETRIEVED_TOPIC, message)
}

func (e *seratEventEmitter) EmitUpdatedEvent(message *Message) error {
	return e.eventEmitter.Publish(SERAT_UPDATED_TOPIC, message)
}

func NewSeratEventEmitter(producer nsq.Producer) SeratEventEmitter {
	return &seratEventEmitter{
		eventEmitter: EventEmitter{
			producer: producer,
		},
	}
}
