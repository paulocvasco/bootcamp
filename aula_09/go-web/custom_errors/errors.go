package customErrors

import "errors"

type CustomError error

func newError(msg string) CustomError {
	return errors.New(msg)
}

var (
	ErrorMissingCode     = newError("field code is required")
	ErrorMissingCoin     = newError("field coin is required")
	ErrorMissingIssuer   = newError("field issuer is required")
	ErrorMissingReciever = newError("field reciever is required")
	ErrorMissingValue    = newError("field value is required")
)
