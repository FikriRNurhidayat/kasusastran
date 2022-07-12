package event

import (
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
)

type WulanganEventEmitter interface {
	EmitCreatedEvent(*message.Wulangan) error
	EmitRetrievedEvent(*message.Wulangan) error
	EmitListedEvent(*message.Wulangans) error
	EmitUpdatedEvent(*message.Wulangan) error
	EmitDeletedEvent(*message.Wulangan) error
}

type wulanganEventEmitter struct {
	eventEmitter EventEmitter
}

func (e *wulanganEventEmitter) EmitCreatedEvent(m *message.Wulangan) error {
	return e.eventEmitter.Publish(WULANGAN_CREATED_TOPIC, m)
}

func (e *wulanganEventEmitter) EmitDeletedEvent(m *message.Wulangan) error {
	return e.eventEmitter.Publish(WULANGAN_DELETED_TOPIC, m)
}

func (e *wulanganEventEmitter) EmitListedEvent(m *message.Wulangans) error {
	return e.eventEmitter.Publish(WULANGAN_LISTED_TOPIC, m)
}

func (e *wulanganEventEmitter) EmitRetrievedEvent(m *message.Wulangan) error {
	return e.eventEmitter.Publish(WULANGAN_RETRIEVED_TOPIC, m)
}

func (e *wulanganEventEmitter) EmitUpdatedEvent(m *message.Wulangan) error {
	return e.eventEmitter.Publish(WULANGAN_UPDATED_TOPIC, m)
}

func NewWulanganEventEmitter(eventEmitter EventEmitter) WulanganEventEmitter {
	return &wulanganEventEmitter{
		eventEmitter: eventEmitter,
	}
}
