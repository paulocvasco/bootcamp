package customErrors

import (
	"errors"
	"fmt"
)

type CustomError error

func newError(msg string) CustomError {
	return errors.New(msg)
}

var (
	ErrorInvalidID        = newError("invalid id")
	ErrorInvalidParameter = newError("id must be a integer")
	ErrorMissingCode      = newError("field code is required")
	ErrorMissingCoin      = newError("field coin is required")
	ErrorMissingIssuer    = newError("field issuer is required")
	ErrorMissingReciever  = newError("field reciever is required")
	ErrorMissingValue     = newError("field value is required")
	ErrorInvalidToken     = newError("invalid authentication token")
)

func WrapError(err, newError error) CustomError {
	return fmt.Errorf("%s : %s", newError, err)
}
