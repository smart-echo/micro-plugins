package grpc

import (
	"net/http"

	perror "github.com/smart-echo/micro/proto/errors/v1"
	"google.golang.org/grpc/codes"
)

var errMapping = map[int32]codes.Code{
	http.StatusOK:                  codes.OK,
	http.StatusBadRequest:          codes.InvalidArgument,
	http.StatusRequestTimeout:      codes.DeadlineExceeded,
	http.StatusNotFound:            codes.NotFound,
	http.StatusConflict:            codes.AlreadyExists,
	http.StatusForbidden:           codes.PermissionDenied,
	http.StatusUnauthorized:        codes.Unauthenticated,
	http.StatusPreconditionFailed:  codes.FailedPrecondition,
	http.StatusNotImplemented:      codes.Unimplemented,
	http.StatusInternalServerError: codes.Internal,
	http.StatusServiceUnavailable:  codes.Unavailable,
}

func microError(err *perror.Error) codes.Code {
	if err == nil {
		return codes.OK
	}

	if code, ok := errMapping[err.Code]; ok {
		return code
	}
	return codes.Unknown
}
