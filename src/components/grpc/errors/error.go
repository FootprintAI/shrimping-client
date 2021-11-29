package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/footprintai/shrimping/components/grpc/flags"
)

func ParseError(err error) *Error {
	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.OK:
			return nil
		case codes.Unauthenticated:
			return newError("Require Login first.", err)
		case codes.Unavailable:
			return newError("Server is unavailable, please retry again later.", err)
		case codes.PermissionDenied:
			return newError("You don't have enough permission to perform such operations.", err)
		default:
			return newError("Unknown", err)
		}
	}
	return newError("not grpc error", err)
}

func newError(plainText string, err error) *Error {
	return &Error{
		PlainText: plainText,
		Details:   err,
	}
}

type Error struct {
	PlainText string
	Details   error
}

func (m *Error) Error() string {
	if flags.ErrorVerbose {
		return fmt.Sprintf("%s(details:%s)", m.PlainText, m.Details.Error())
	}
	return m.PlainText
}
