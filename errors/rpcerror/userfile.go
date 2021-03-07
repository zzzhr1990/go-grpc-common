package rpcerror

import (

	//"time"

	// pb "github.com/zzzhr1990/common-grpc/go/services/offline/systemtask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// log "github.com/sirupsen/logrus"
)

var (
	// ErrorIllegalFilePathIdentity exc
	ErrorIllegalFilePathIdentity = status.Errorf(codes.InvalidArgument, "ILLEGAL_FILE_PATH_OR_IDENTITY")
	// ErrInternalError interbal, cause by unknown reason
	ErrInternalError = status.Errorf(codes.Internal, "INTERNAL_ERROR")
	// ErrFileExists a file exists and not allow to override it
	ErrFileExists = status.Errorf(codes.AlreadyExists, "FILE_ALREADY_EXISTS")
	/*
		// ErrCannotAquireFileLock cannot get lock
		ErrCannotAquireFileLock = status.Errorf(codes.FailedPrecondition, "CANNOT_AQUIRE_FILE_LOCK")
		// FileIsLocking file is locking, and you cannot do any write operation
		FileIsLocking = status.Errorf(codes.FailedPrecondition, "FILE_IS_LOCKING")
	*/

	// ErrParentContainsFile parent has file, Loop operation detected
	ErrParentContainsFile = status.Errorf(codes.FailedPrecondition, "PARENT_CONTAINS_FILE")
	// ErrEditRootDirectory you're trying to operate root directory
	ErrEditRootDirectory = status.Errorf(codes.FailedPrecondition, "EDIT_ROOT_DIR")
	// ErrCopyRootDirectory parent has file
	ErrCopyRootDirectory = status.Errorf(codes.FailedPrecondition, "COPY_ROOT_DIR")
	// OperationContainsSelf parent has file
	OperationContainsSelf = status.Errorf(codes.FailedPrecondition, "OPERATION_CONTAINS_SELF")
	// ErrStoreFileNotFound parent has file
	ErrStoreFileNotFound = status.Errorf(codes.NotFound, "STORE_FILE_NOT_FOUND")
	// ErrUserIdentityIsEmpty parent has file
	ErrUserIdentityIsEmpty = status.Errorf(codes.FailedPrecondition, "USER_IDENTITY_EMPTY")
	// ErrParentDirectoryNotFound parent has file
	ErrParentDirectoryNotFound = status.Errorf(codes.NotFound, "PARENT_DIRECTORY_NOT_FOUND")
	// ErrSourceFileNotFound a source for copy... not found.
	ErrSourceFileNotFound = status.Errorf(codes.NotFound, "SOURCE_FILE_NOT_FOUND")
	// ErrLoopOverflow a source cause too many loop
	ErrLoopOverflow = status.Errorf(codes.FailedPrecondition, "LOOP_OVERFLOW")

	// ErrMissingFilename a source cause too many loop
	ErrMissingFilename = status.Errorf(codes.FailedPrecondition, "MISSING_FILENAME")
	// ErrTaskCancel a source cause too many loop
	ErrTaskCancel = status.Errorf(codes.FailedPrecondition, "TASK_CANCEL")

	// ErrIllegalHash empty identity
	ErrIllegalHash = status.Errorf(codes.InvalidArgument, "ILLEGAL_HASH")

	// ErrMissingProps empty identity
	ErrMissingProps = status.Errorf(codes.InvalidArgument, "MISSING_PROPS")

	// ErrMissingUserIdentity miss uid
	ErrMissingUserIdentity = status.Errorf(codes.InvalidArgument, "MISSING_USER_IDENTITY")

	// ErrIllegalFile file illegal
	ErrIllegalFile = status.Errorf(codes.FailedPrecondition, "ILLEGAL_FILE")

	// ErrPathLocked empty identity
	ErrPathLocked = status.Errorf(codes.FailedPrecondition, "PATH_LOCKED")

	// TaskAlreadyExists empty identity
	TaskAlreadyExists = status.Errorf(codes.FailedPrecondition, "TASK_ALREADY_EXIST")
)
