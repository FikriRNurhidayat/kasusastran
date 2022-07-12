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

type MockSeratEventEmitter struct {
	eventEmitter *mocks.EventEmitter
}

func TestSeratEventEmitter_EmitCreatedEvent(t *testing.T) {
	type input struct {
		message *message.Serat
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratEventEmitter, *input, *output)
	}

	tests := []scenario{
		{
			name: "OK",
			in: &input{
				message: &message.Serat{
					ID:        uuid.New(),
					Kind:      event.SERAT_CREATED_TOPIC,
					CreatedAt: time.Now(),
					Actor:     &message.Actor{},
					Payload: &entity.Serat{
						ID:                uuid.New(),
						Title:             "Lorem ipsum dolor sit amet",
						Description:       "Lorem ipsum dolor sit amet",
						Content:           "Lorem ipsum dolor sit amet",
						CoverImageUrl:     "http://placekitten.com/192/108",
						ThumbnailImageUrl: "http://placekitten.com/192/108",
					},
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockSeratEventEmitter, in *input, out *output) {
				m.eventEmitter.On("Publish", event.SERAT_CREATED_TOPIC, in.message).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratEventEmitter{
				eventEmitter: new(mocks.EventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := event.NewSeratEventEmitter(m.eventEmitter)
			err := subject.EmitCreatedEvent(tt.in.message)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}

func TestSeratEventEmitter_EmitDeletedEvent(t *testing.T) {
	type input struct {
		message *message.Serat
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratEventEmitter, *input, *output)
	}

	tests := []scenario{
		{
			name: "OK",
			in: &input{
				message: &message.Serat{
					ID:        uuid.New(),
					Kind:      event.SERAT_DELETED_TOPIC,
					CreatedAt: time.Now(),
					Actor:     &message.Actor{},
					Payload: &entity.Serat{
						ID:                uuid.New(),
						Title:             "Lorem ipsum dolor sit amet",
						Description:       "Lorem ipsum dolor sit amet",
						Content:           "Lorem ipsum dolor sit amet",
						CoverImageUrl:     "http://placekitten.com/192/108",
						ThumbnailImageUrl: "http://placekitten.com/192/108",
					},
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockSeratEventEmitter, in *input, out *output) {
				m.eventEmitter.On("Publish", event.SERAT_DELETED_TOPIC, in.message).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratEventEmitter{
				eventEmitter: new(mocks.EventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := event.NewSeratEventEmitter(m.eventEmitter)
			err := subject.EmitDeletedEvent(tt.in.message)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}

func TestSeratEventEmitter_EmitUpdatedEvent(t *testing.T) {
	type input struct {
		message *message.Serat
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratEventEmitter, *input, *output)
	}

	tests := []scenario{
		{
			name: "OK",
			in: &input{
				message: &message.Serat{
					ID:        uuid.New(),
					Kind:      event.SERAT_UPDATED_TOPIC,
					CreatedAt: time.Now(),
					Actor:     &message.Actor{},
					Payload: &entity.Serat{
						ID:                uuid.New(),
						Title:             "Lorem ipsum dolor sit amet",
						Description:       "Lorem ipsum dolor sit amet",
						Content:           "Lorem ipsum dolor sit amet",
						CoverImageUrl:     "http://placekitten.com/192/108",
						ThumbnailImageUrl: "http://placekitten.com/192/108",
					},
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockSeratEventEmitter, in *input, out *output) {
				m.eventEmitter.On("Publish", event.SERAT_UPDATED_TOPIC, in.message).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratEventEmitter{
				eventEmitter: new(mocks.EventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := event.NewSeratEventEmitter(m.eventEmitter)
			err := subject.EmitUpdatedEvent(tt.in.message)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}

func TestSeratEventEmitter_EmitRetrievedEvent(t *testing.T) {
	type input struct {
		message *message.Serat
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratEventEmitter, *input, *output)
	}

	tests := []scenario{
		{
			name: "OK",
			in: &input{
				message: &message.Serat{
					ID:        uuid.New(),
					Kind:      event.SERAT_RETRIEVED_TOPIC,
					CreatedAt: time.Now(),
					Actor:     &message.Actor{},
					Payload: &entity.Serat{
						ID:                uuid.New(),
						Title:             "Lorem ipsum dolor sit amet",
						Description:       "Lorem ipsum dolor sit amet",
						Content:           "Lorem ipsum dolor sit amet",
						CoverImageUrl:     "http://placekitten.com/192/108",
						ThumbnailImageUrl: "http://placekitten.com/192/108",
					},
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockSeratEventEmitter, in *input, out *output) {
				m.eventEmitter.On("Publish", event.SERAT_RETRIEVED_TOPIC, in.message).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratEventEmitter{
				eventEmitter: new(mocks.EventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := event.NewSeratEventEmitter(m.eventEmitter)
			err := subject.EmitRetrievedEvent(tt.in.message)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}

func TestSeratEventEmitter_EmitListedEvent(t *testing.T) {
	type input struct {
		message *message.Serats
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratEventEmitter, *input, *output)
	}

	tests := []scenario{
		{
			name: "OK",
			in: &input{
				message: &message.Serats{
					ID:        uuid.New(),
					Kind:      event.SERAT_LISTED_TOPIC,
					CreatedAt: time.Now(),
					Actor:     &message.Actor{},
					Payload: []*entity.Serat{
						{
							ID:                uuid.New(),
							Title:             "Lorem ipsum dolor sit amet",
							Description:       "Lorem ipsum dolor sit amet",
							Content:           "Lorem ipsum dolor sit amet",
							CoverImageUrl:     "http://placekitten.com/192/108",
							ThumbnailImageUrl: "http://placekitten.com/192/108",
						},
					},
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockSeratEventEmitter, in *input, out *output) {
				m.eventEmitter.On("Publish", event.SERAT_LISTED_TOPIC, in.message).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratEventEmitter{
				eventEmitter: new(mocks.EventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := event.NewSeratEventEmitter(m.eventEmitter)
			err := subject.EmitListedEvent(tt.in.message)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}
