package event_test

import (
	"testing"
	"time"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	mocks "github.com/fikrirnurhidayat/kasusastran/mocks/domain/event"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockUserEventEmitter struct {
	eventEmitter *mocks.EventEmitter
}

func TestUserEventEmitter_EmitRegisteredEvent(t *testing.T) {
	type input struct {
		message *message.User
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockUserEventEmitter, *input, *output)
	}

	tests := []scenario{
		{
			name: "OK",
			in: &input{
				message: &message.User{
					ID:        uuid.New(),
					Kind:      event.USER_REGISTERED_TOPIC,
					CreatedAt: time.Now(),
					Actor:     &message.Actor{},
					Payload: &entity.User{
						ID:                uuid.New(),
						Name:              "Fikri",
						Email:             "fikrirnurhidayat@gmail.com",
						EncryptedPassword: "$2a$10$O2pp2NvX/Y3OgBM66NUrjOtASWhg3rMhft4X0Ii4U8gX3AOJqcItK",
						CreatedAt:         time.Now(),
						UpdatedAt:         time.Now(),
					},
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockUserEventEmitter, in *input, out *output) {
				in.message.Actor.ID = in.message.Payload.ID.String()
				in.message.Actor.Kind = "USER"

				m.eventEmitter.On("Publish", event.USER_REGISTERED_TOPIC, in.message).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockUserEventEmitter{
				eventEmitter: new(mocks.EventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := event.NewUserEventEmitter(m.eventEmitter)
			err := subject.EmitRegisteredEvent(tt.in.message)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}
