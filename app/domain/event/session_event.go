package event

import "github.com/fikrirnurhidayat/kasusastran/app/domain/message"

type SessionEventEmitter interface {
	EmitCreatedEvent(*message.User) error
}

type sessionEventEmitter struct {
	eventEmitter EventEmitter
}

// EmitCreatedEvent implements SessionEventEmitter
func (e *sessionEventEmitter) EmitCreatedEvent(m *message.User) error {
	return e.eventEmitter.Publish(SESSION_CREATED_TOPIC, m)
}

func NewSessionEventEmitter(eventEmitter EventEmitter) SessionEventEmitter {
	return &sessionEventEmitter{
		eventEmitter: eventEmitter,
	}
}
