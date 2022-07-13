package listener

import (
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	"github.com/nsqio/go-nsq"
	"google.golang.org/grpc/grpclog"
)

type SeratListener interface {
	CreatedEventListener(message *nsq.Message) error
	RetrievedEventListener(message *nsq.Message) error
	ListedEventListener(message *nsq.Message) error
	UpdatedEventListener(message *nsq.Message) error
	DeletedEventListener(message *nsq.Message) error
}

type seratListener struct {
	listener Listener
	logger   grpclog.LoggerV2
}

func NewSeratListener(
	logger grpclog.LoggerV2,
	listener Listener,
) SeratListener {
	return &seratListener{
		listener: listener,
		logger:   logger,
	}
}

func (s *seratListener) CreatedEventListener(m *nsq.Message) (err error) {
	var createdSeratEventPayload message.Serat

	err = s.listener.Parse(m.Body, &createdSeratEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", createdSeratEventPayload)

	return nil
}

func (s *seratListener) ListedEventListener(m *nsq.Message) (err error) {
	var listedSeratEventPayload message.Serats

	err = s.listener.Parse(m.Body, &listedSeratEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", listedSeratEventPayload)

	return nil
}

func (s *seratListener) RetrievedEventListener(m *nsq.Message) (err error) {
	var retrievedSeratEventPayload message.Serat

	err = s.listener.Parse(m.Body, &retrievedSeratEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", retrievedSeratEventPayload)

	return nil
}

func (s *seratListener) UpdatedEventListener(m *nsq.Message) (err error) {
	var updatedSeratEventPayload message.Serat

	err = s.listener.Parse(m.Body, &updatedSeratEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", updatedSeratEventPayload)

	return nil
}

func (s *seratListener) DeletedEventListener(m *nsq.Message) (err error) {
	var deletedSeratEventPayload message.Serat

	err = s.listener.Parse(m.Body, &deletedSeratEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", deletedSeratEventPayload)

	return nil
}
