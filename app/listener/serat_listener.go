package listener

import (
	"github.com/fikrirnurhidayat/kasusastran/app/domain/message"
	"github.com/nsqio/go-nsq"
	"google.golang.org/grpc/grpclog"
)

type SeratListener interface {
	CreatedEventListener(message *nsq.Message) error
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

	return nil
}
