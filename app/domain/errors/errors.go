package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidUUID = status.Error(codes.InvalidArgument, "invalid uuid")
)
