package user

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	//ErrorNameExist when user exists
	ErrorNameExist = status.New(codes.AlreadyExists, "USER_NAME_EXISTS").Err()
	//ErrorPhoneExist when phone exists
	ErrorPhoneExist = status.New(codes.AlreadyExists, "USER_PHONE_EXISTS").Err()
	//ErrorLoginFailed login in failed
	ErrorLoginFailed = status.New(codes.Unauthenticated, "LOGIN_FAILED").Err()
	//ErrorUserNotFound user not exists
	ErrorUserNotFound = status.New(codes.NotFound, "USER_NOT_FOUND").Err()
	//ErrorPasswordInvalid login in failed
	ErrorPasswordInvalid = status.New(codes.InvalidArgument, "PASSWORD_INVALID").Err()
)
