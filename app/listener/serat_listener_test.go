package listener_test

import (
	"errors"
	"testing"

	"github.com/nsqio/go-nsq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/fikrirnurhidayat/kasusastran/app/listener"
	mockListener "github.com/fikrirnurhidayat/kasusastran/mocks/listener"
	mockPackage "github.com/fikrirnurhidayat/kasusastran/mocks/package"
)

type MockSeratListener struct {
	listener *mockListener.Listener
	logger   *mockPackage.LoggerV2
}

func TestSeratListener_CreatedEventListener(t *testing.T) {
	type input struct {
		m *nsq.Message
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratListener, *input, *output)
	}

	tests := []scenario{
		{
			name: "s.listener.Parse: failed to parse message",
			in: &input{
				m: &nsq.Message{},
			},
			out: &output{
				err: errors.New("s.listener.Parse: failed to parse message"),
			},
			on: func(m *MockSeratListener, i *input, o *output) {
				m.listener.On("Parse", i.m.Body, mock.AnythingOfType("*message.Serat")).Return(o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				m: &nsq.Message{},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockSeratListener, i *input, o *output) {
				m.listener.On("Parse", i.m.Body, mock.AnythingOfType("*message.Serat")).Return(nil)
				m.logger.On("Infof", "Received: %v", mock.AnythingOfType("message.Serat"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratListener{
				listener: new(mockListener.Listener),
				logger:   new(mockPackage.LoggerV2),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := listener.NewSeratListener(m.logger, m.listener)
			err := subject.CreatedEventListener(tt.in.m)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSeratListener_UpdatedEventListener(t *testing.T) {
	type input struct {
		m *nsq.Message
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratListener, *input, *output)
	}

	tests := []scenario{
		{
			name: "s.listener.Parse: failed to parse message",
			in: &input{
				m: &nsq.Message{},
			},
			out: &output{
				err: errors.New("s.listener.Parse: failed to parse message"),
			},
			on: func(m *MockSeratListener, i *input, o *output) {
				m.listener.On("Parse", i.m.Body, mock.AnythingOfType("*message.Serat")).Return(o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				m: &nsq.Message{},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockSeratListener, i *input, o *output) {
				m.listener.On("Parse", i.m.Body, mock.AnythingOfType("*message.Serat")).Return(nil)
				m.logger.On("Infof", "Received: %v", mock.AnythingOfType("message.Serat"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratListener{
				listener: new(mockListener.Listener),
				logger:   new(mockPackage.LoggerV2),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := listener.NewSeratListener(m.logger, m.listener)
			err := subject.UpdatedEventListener(tt.in.m)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSeratListener_RetrievedEventListener(t *testing.T) {
	type input struct {
		m *nsq.Message
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratListener, *input, *output)
	}

	tests := []scenario{
		{
			name: "s.listener.Parse: failed to parse message",
			in: &input{
				m: &nsq.Message{},
			},
			out: &output{
				err: errors.New("s.listener.Parse: failed to parse message"),
			},
			on: func(m *MockSeratListener, i *input, o *output) {
				m.listener.On("Parse", i.m.Body, mock.AnythingOfType("*message.Serat")).Return(o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				m: &nsq.Message{},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockSeratListener, i *input, o *output) {
				m.listener.On("Parse", i.m.Body, mock.AnythingOfType("*message.Serat")).Return(nil)
				m.logger.On("Infof", "Received: %v", mock.AnythingOfType("message.Serat"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratListener{
				listener: new(mockListener.Listener),
				logger:   new(mockPackage.LoggerV2),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := listener.NewSeratListener(m.logger, m.listener)
			err := subject.RetrievedEventListener(tt.in.m)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSeratListener_ListedEventListener(t *testing.T) {
	type input struct {
		m *nsq.Message
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratListener, *input, *output)
	}

	tests := []scenario{
		{
			name: "s.listener.Parse: failed to parse message",
			in: &input{
				m: &nsq.Message{},
			},
			out: &output{
				err: errors.New("s.listener.Parse: failed to parse message"),
			},
			on: func(m *MockSeratListener, i *input, o *output) {
				m.listener.On("Parse", i.m.Body, mock.AnythingOfType("*message.Serats")).Return(o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				m: &nsq.Message{},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockSeratListener, i *input, o *output) {
				m.listener.On("Parse", i.m.Body, mock.AnythingOfType("*message.Serats")).Return(nil)
				m.logger.On("Infof", "Received: %v", mock.AnythingOfType("message.Serats"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratListener{
				listener: new(mockListener.Listener),
				logger:   new(mockPackage.LoggerV2),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := listener.NewSeratListener(m.logger, m.listener)
			err := subject.ListedEventListener(tt.in.m)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSeratListener_DeletedEventListener(t *testing.T) {
	type input struct {
		m *nsq.Message
	}

	type output struct {
		err error
	}

	type scenario struct {
		name string
		in   *input
		out  *output
		on   func(*MockSeratListener, *input, *output)
	}

	tests := []scenario{
		{
			name: "s.listener.Parse: failed to parse message",
			in: &input{
				m: &nsq.Message{},
			},
			out: &output{
				err: errors.New("s.listener.Parse: failed to parse message"),
			},
			on: func(m *MockSeratListener, i *input, o *output) {
				m.listener.On("Parse", i.m.Body, mock.AnythingOfType("*message.Serat")).Return(o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				m: &nsq.Message{},
			},
			out: &output{
				err: nil,
			},
			on: func(m *MockSeratListener, i *input, o *output) {
				m.listener.On("Parse", i.m.Body, mock.AnythingOfType("*message.Serat")).Return(nil)
				m.logger.On("Infof", "Received: %v", mock.AnythingOfType("message.Serat"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockSeratListener{
				listener: new(mockListener.Listener),
				logger:   new(mockPackage.LoggerV2),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := listener.NewSeratListener(m.logger, m.listener)
			err := subject.DeletedEventListener(tt.in.m)

			if tt.out.err != nil {
				assert.NotNil(t, err)
				assert.Contains(t, tt.out.err.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
