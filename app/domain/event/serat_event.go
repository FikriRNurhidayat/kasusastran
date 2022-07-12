package event

import (
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
)

type SeratEventEmitter interface {
	EmitCreatedEvent(*message.Serat) error
	EmitRetrievedEvent(*message.Serat) error
	EmitListedEvent(*message.Serats) error
	EmitUpdatedEvent(*message.Serat) error
	EmitDeletedEvent(*message.Serat) error
}

type seratEventEmitter struct {
	eventEmitter EventEmitter
}

func (e *seratEventEmitter) EmitCreatedEvent(m *message.Serat) error {
	return e.eventEmitter.Publish(SERAT_CREATED_TOPIC, m)
}

func (e *seratEventEmitter) EmitDeletedEvent(m *message.Serat) error {
	return e.eventEmitter.Publish(SERAT_DELETED_TOPIC, m)
}

func (e *seratEventEmitter) EmitListedEvent(m *message.Serats) error {
	return e.eventEmitter.Publish(SERAT_LISTED_TOPIC, m)
}

func (e *seratEventEmitter) EmitRetrievedEvent(m *message.Serat) error {
	return e.eventEmitter.Publish(SERAT_RETRIEVED_TOPIC, m)
}

func (e *seratEventEmitter) EmitUpdatedEvent(m *message.Serat) error {
	return e.eventEmitter.Publish(SERAT_UPDATED_TOPIC, m)
}

func NewSeratEventEmitter(eventEmitter EventEmitter) SeratEventEmitter {
	return &seratEventEmitter{
		eventEmitter: eventEmitter,
	}
}
