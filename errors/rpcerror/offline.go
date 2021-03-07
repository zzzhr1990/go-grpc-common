package rpcerror

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrEmptyTaskIdentity empty identity
	ErrEmptyTaskIdentity = status.Errorf(codes.InvalidArgument, "EMPTY_TASK_IDENTITY")
	// ErrEmptyUserIdentity empty identity
	ErrEmptyUserIdentity = status.Errorf(codes.InvalidArgument, "EMPTY_USER_IDENTITY")
	// ErrEmptyFileHash empty hash
	ErrEmptyFileHash = status.Errorf(codes.InvalidArgument, "EMPTY_FILE_HASH")
	// ErrEmptyRequest empty identity
	ErrEmptyRequest = status.Errorf(codes.InvalidArgument, "EMPTY_REQUEST")
	// ErrUnsupportDownload empty identity
	ErrUnsupportDownload = status.Errorf(codes.InvalidArgument, "UNSUPPORT_DOWNLOAD")
	// ErrEmptyPath empty identity
	ErrEmptyPath = status.Errorf(codes.InvalidArgument, "EMPTY_PATH")

	// ErrTaskBaseLocked empty identity
	ErrTaskBaseLocked = status.Errorf(codes.FailedPrecondition, "TASK_BASE_LOCKED")
)
