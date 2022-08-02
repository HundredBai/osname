package osname

import (
	"errors"
	"fmt"
)

type err string

func (e err) Error() string {
	return string(e)
}

func (e err) Cause(cause error) error {
	return &withCause{
		k:     e,
		cause: cause,
	}
}

type withCause struct {
	k     error
	cause error
}

func (w *withCause) Error() string {
	return fmt.Sprintf("%s: %s", w.k.Error(), w.cause.Error())
}

func (w *withCause) Unwrap() error {
	return w.cause
}

func (w *withCause) Is(target error) bool {
	return errors.Is(w.k, target) || errors.Is(w.cause, target)
}
