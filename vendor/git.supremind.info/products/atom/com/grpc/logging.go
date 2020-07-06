package grpc

import (
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
)

// assign log level for each grpc response code
var codeMappings = map[codes.Code]zapcore.Level{
	codes.OK:                 zapcore.InfoLevel,
	codes.Canceled:           zapcore.WarnLevel,
	codes.Unknown:            zapcore.WarnLevel,
	codes.InvalidArgument:    zapcore.WarnLevel,
	codes.DeadlineExceeded:   zapcore.WarnLevel,
	codes.NotFound:           zapcore.WarnLevel,
	codes.PermissionDenied:   zapcore.WarnLevel,
	codes.ResourceExhausted:  zapcore.WarnLevel,
	codes.FailedPrecondition: zapcore.WarnLevel,
	codes.Aborted:            zapcore.WarnLevel,
	codes.OutOfRange:         zapcore.WarnLevel,
	codes.Unimplemented:      zapcore.ErrorLevel,
	codes.Internal:           zapcore.ErrorLevel,
	codes.Unavailable:        zapcore.ErrorLevel,
	codes.DataLoss:           zapcore.ErrorLevel,
	codes.Unauthenticated:    zapcore.WarnLevel,
}

func CodeToLevel(code codes.Code) zapcore.Level {
	return codeMappings[code]
}
