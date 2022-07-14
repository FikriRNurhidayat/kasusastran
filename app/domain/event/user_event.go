package event

import "github.com/fikrirnurhidayat/kasusastran/app/domain/message"

type UserEventEmitter interface {
	EmitRegisteredEvent(*message.User) error
}

type userEventEmitter struct {
	eventEmitter EventEmitter
}

func (e *userEventEmitter) EmitRegisteredEvent(m *message.User) error {
	return e.eventEmitter.Publish(USER_REGISTERED_TOPIC, m)
}

func NewUserEventEmitter(eventEmitter EventEmitter) UserEventEmitter {
	return &userEventEmitter{
		eventEmitter: eventEmitter,
	}
}
