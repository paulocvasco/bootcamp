package customErrors

import "errors"

type CustomError error

func newError(msg string) CustomError {
	return errors.New(msg)
}

var (
	ErrorMissingCode     = newError("field code is obrigatory")
	ErrorMissingCoin     = newError("field coin is obrigatory")
	ErrorMissingIssuer   = newError("field issuer is obrigatory")
	ErrorMissingReciever = newError("field reciever is obrigatory")
	ErrorMissingValue    = newError("field value is obrigatory")
)
