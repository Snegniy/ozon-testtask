package apperror

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound = errors.New("link not found")
	ErrIsEmpty  = errors.New("url cannot be empty")

	ErrGrpcIsEmpty     = status.Error(codes.InvalidArgument, "url cannot be empty")
	ErrGrpcNotFound    = status.Error(codes.NotFound, "link not found")
	ErrGrpcServerError = status.Error(codes.Internal, "internal server error")
)
