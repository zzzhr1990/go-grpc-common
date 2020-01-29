package taskerror

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrInternal When file in cold storage
	ErrInternal = status.Errorf(codes.Internal, "INTERNAL_ERROR")
)
