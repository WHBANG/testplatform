package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IsAlreadyExists tests AlreadyExists error
func IsAlreadyExists(e error) bool {
	s, ok := status.FromError(e)
	return ok && s.Code() == codes.AlreadyExists
}

// IsNotFound tests NotFound error
func IsNotFound(e error) bool {
	s, ok := status.FromError(e)
	return ok && s.Code() == codes.NotFound
}

// IsInvalidArgument tests InvalidArgument error
func IsInvalidArgument(e error) bool {
	s, ok := status.FromError(e)
	return ok && s.Code() == codes.InvalidArgument
}

func IsUnauthenticated(e error) bool {
	s, ok := status.FromError(e)
	return ok && s.Code() == codes.Unauthenticated
}

func IsRetryableError(e error) bool {
	s, ok := status.FromError(e)
	return ok && (s.Code() == codes.Unauthenticated) // client will try to do auth for user if auth failed
	// TODO: enable this when alluxio server returns correct codes
	//codes.Unavailable ||
	//s.Code() == codes.FailedPrecondition ||
	//s.Code() == codes.Aborted)
}
