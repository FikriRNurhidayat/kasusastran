package srv

import "google.golang.org/grpc/grpclog"

type Server interface {
	setLogger(grpclog.LoggerV2)
}

func WithLogger[T Server](logger grpclog.LoggerV2) func(T) {
	return func(s T) {
		s.setLogger(logger)
	}
}
