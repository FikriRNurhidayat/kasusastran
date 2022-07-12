package listener

import (
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	"github.com/nsqio/go-nsq"
	"google.golang.org/grpc/grpclog"
)

type SeratListener interface {
	CreatedEventListener(message *nsq.Message) error
	RetrievedEventListener(message *nsq.Message) error
	UpdatedEventListener(message *nsq.Message) error
	DeletedEventListener(message *nsq.Message) error
}

type seratListener struct {
	logger grpclog.LoggerV2
}

func NewSeratListener(
	logger grpclog.LoggerV2,
) SeratListener {
	return &seratListener{
		logger: logger,
	}
}

func (s *seratListener) CreatedEventListener(m *nsq.Message) (err error) {
	var createdSeratEventPayload *message.Serat

	err = Parse(m.Body, &createdSeratEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", createdSeratEventPayload)

	return nil
}

func (s *seratListener) RetrievedEventListener(m *nsq.Message) (err error) {
	var retrievedSeratEventPayload *message.Serat

	err = Parse(m.Body, &retrievedSeratEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", retrievedSeratEventPayload)

	return nil
}

func (s *seratListener) UpdatedEventListener(m *nsq.Message) (err error) {
	var updatedSeratEventPayload *message.Serat

	err = Parse(m.Body, &updatedSeratEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", updatedSeratEventPayload)

	return nil
}

func (s *seratListener) DeletedEventListener(m *nsq.Message) (err error) {
	var deletedSeratEventPayload *message.Serat

	err = Parse(m.Body, &deletedSeratEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", deletedSeratEventPayload)

	return nil
}
