package listener

import (
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	"github.com/nsqio/go-nsq"
	"google.golang.org/grpc/grpclog"
)

type WulanganListener interface {
	CreatedEventListener(message *nsq.Message) error
	RetrievedEventListener(message *nsq.Message) error
	ListedEventListener(message *nsq.Message) error
	UpdatedEventListener(message *nsq.Message) error
	DeletedEventListener(message *nsq.Message) error
}

type wulanganListener struct {
	listener Listener
	logger   grpclog.LoggerV2
}

func NewWulanganListener(
	logger grpclog.LoggerV2,
	listener Listener,
) WulanganListener {
	return &wulanganListener{
		listener: listener,
		logger:   logger,
	}
}

func (s *wulanganListener) CreatedEventListener(m *nsq.Message) (err error) {
	var createdWulanganEventPayload message.Wulangan

	err = s.listener.Parse(m.Body, &createdWulanganEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", createdWulanganEventPayload)

	return nil
}

func (s *wulanganListener) ListedEventListener(m *nsq.Message) (err error) {
	var listedWulanganEventPayload message.Wulangans

	err = s.listener.Parse(m.Body, &listedWulanganEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", listedWulanganEventPayload)

	return nil
}

func (s *wulanganListener) RetrievedEventListener(m *nsq.Message) (err error) {
	var retrievedWulanganEventPayload message.Wulangan

	err = s.listener.Parse(m.Body, &retrievedWulanganEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", retrievedWulanganEventPayload)

	return nil
}

func (s *wulanganListener) UpdatedEventListener(m *nsq.Message) (err error) {
	var updatedWulanganEventPayload message.Wulangan

	err = s.listener.Parse(m.Body, &updatedWulanganEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", updatedWulanganEventPayload)

	return nil
}

func (s *wulanganListener) DeletedEventListener(m *nsq.Message) (err error) {
	var deletedWulanganEventPayload message.Wulangan

	err = s.listener.Parse(m.Body, &deletedWulanganEventPayload)
	if err != nil {
		return err
	}

	s.logger.Infof("Received: %v", deletedWulanganEventPayload)

	return nil
}
