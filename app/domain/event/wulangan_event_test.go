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

type MockWulanganEventEmitter struct {
	eventEmitter *mocks.EventEmitter
}

func TestWulanganEventEmitter_EmitCreatedEvent(t *testing.T) {
	type input struct {
		message *message.Wulangan
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockWulanganEventEmitter, *input, *output)
	}

	tests := []scenario{
		{
			name: "OK",
			in: &input{
				message: &message.Wulangan{
					ID:        uuid.New(),
					Kind:      event.WULANGAN_CREATED_TOPIC,
					CreatedAt: time.Now(),
					Actor:     &message.Actor{},
					Payload: &entity.Wulangan{
						ID:                uuid.New(),
						Title:             "Lorem ipsum dolor sit amet",
						Description:       "Lorem ipsum dolor sit amet",
						CoverImageUrl:     "http://placekitten.com/192/108",
						ThumbnailImageUrl: "http://placekitten.com/192/108",
					},
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockWulanganEventEmitter, in *input, out *output) {
				m.eventEmitter.On("Publish", event.WULANGAN_CREATED_TOPIC, in.message).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockWulanganEventEmitter{
				eventEmitter: new(mocks.EventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := event.NewWulanganEventEmitter(m.eventEmitter)
			err := subject.EmitCreatedEvent(tt.in.message)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}

func TestWulanganEventEmitter_EmitDeletedEvent(t *testing.T) {
	type input struct {
		message *message.Wulangan
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockWulanganEventEmitter, *input, *output)
	}

	tests := []scenario{
		{
			name: "OK",
			in: &input{
				message: &message.Wulangan{
					ID:        uuid.New(),
					Kind:      event.WULANGAN_DELETED_TOPIC,
					CreatedAt: time.Now(),
					Actor:     &message.Actor{},
					Payload: &entity.Wulangan{
						ID:                uuid.New(),
						Title:             "Lorem ipsum dolor sit amet",
						Description:       "Lorem ipsum dolor sit amet",
						CoverImageUrl:     "http://placekitten.com/192/108",
						ThumbnailImageUrl: "http://placekitten.com/192/108",
					},
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockWulanganEventEmitter, in *input, out *output) {
				m.eventEmitter.On("Publish", event.WULANGAN_DELETED_TOPIC, in.message).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockWulanganEventEmitter{
				eventEmitter: new(mocks.EventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := event.NewWulanganEventEmitter(m.eventEmitter)
			err := subject.EmitDeletedEvent(tt.in.message)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}

func TestWulanganEventEmitter_EmitUpdatedEvent(t *testing.T) {
	type input struct {
		message *message.Wulangan
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockWulanganEventEmitter, *input, *output)
	}

	tests := []scenario{
		{
			name: "OK",
			in: &input{
				message: &message.Wulangan{
					ID:        uuid.New(),
					Kind:      event.WULANGAN_UPDATED_TOPIC,
					CreatedAt: time.Now(),
					Actor:     &message.Actor{},
					Payload: &entity.Wulangan{
						ID:                uuid.New(),
						Title:             "Lorem ipsum dolor sit amet",
						Description:       "Lorem ipsum dolor sit amet",
						CoverImageUrl:     "http://placekitten.com/192/108",
						ThumbnailImageUrl: "http://placekitten.com/192/108",
					},
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockWulanganEventEmitter, in *input, out *output) {
				m.eventEmitter.On("Publish", event.WULANGAN_UPDATED_TOPIC, in.message).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockWulanganEventEmitter{
				eventEmitter: new(mocks.EventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := event.NewWulanganEventEmitter(m.eventEmitter)
			err := subject.EmitUpdatedEvent(tt.in.message)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}

func TestWulanganEventEmitter_EmitRetrievedEvent(t *testing.T) {
	type input struct {
		message *message.Wulangan
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockWulanganEventEmitter, *input, *output)
	}

	tests := []scenario{
		{
			name: "OK",
			in: &input{
				message: &message.Wulangan{
					ID:        uuid.New(),
					Kind:      event.WULANGAN_RETRIEVED_TOPIC,
					CreatedAt: time.Now(),
					Actor:     &message.Actor{},
					Payload: &entity.Wulangan{
						ID:                uuid.New(),
						Title:             "Lorem ipsum dolor sit amet",
						Description:       "Lorem ipsum dolor sit amet",
						CoverImageUrl:     "http://placekitten.com/192/108",
						ThumbnailImageUrl: "http://placekitten.com/192/108",
					},
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockWulanganEventEmitter, in *input, out *output) {
				m.eventEmitter.On("Publish", event.WULANGAN_RETRIEVED_TOPIC, in.message).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockWulanganEventEmitter{
				eventEmitter: new(mocks.EventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := event.NewWulanganEventEmitter(m.eventEmitter)
			err := subject.EmitRetrievedEvent(tt.in.message)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}

func TestWulanganEventEmitter_EmitListedEvent(t *testing.T) {
	type input struct {
		message *message.Wulangans
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockWulanganEventEmitter, *input, *output)
	}

	tests := []scenario{
		{
			name: "OK",
			in: &input{
				message: &message.Wulangans{
					ID:        uuid.New(),
					Kind:      event.WULANGAN_LISTED_TOPIC,
					CreatedAt: time.Now(),
					Actor:     &message.Actor{},
					Payload: []*entity.Wulangan{
						{
							ID:                uuid.New(),
							Title:             "Lorem ipsum dolor sit amet",
							Description:       "Lorem ipsum dolor sit amet",
							CoverImageUrl:     "http://placekitten.com/192/108",
							ThumbnailImageUrl: "http://placekitten.com/192/108",
						},
					},
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockWulanganEventEmitter, in *input, out *output) {
				m.eventEmitter.On("Publish", event.WULANGAN_LISTED_TOPIC, in.message).Return(out.err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockWulanganEventEmitter{
				eventEmitter: new(mocks.EventEmitter),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := event.NewWulanganEventEmitter(m.eventEmitter)
			err := subject.EmitListedEvent(tt.in.message)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			}
		})
	}
}
