package taskerror

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrFileNotReady When file in cold storage
	ErrFileNotReady = status.Errorf(codes.InvalidArgument, "FILE_NOT_READY")
	// ErrTaskIdentityEmpty When recv empty
	ErrTaskIdentityEmpty = status.Errorf(codes.InvalidArgument, "TASK_IDENTITY_EMPTY")
	// ErrTaskNotFound When task not found
	ErrTaskNotFound = status.Errorf(codes.NotFound, "TASK_NOT_FOUND")
)
