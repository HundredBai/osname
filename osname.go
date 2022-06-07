package osname

import (
	"errors"
	"fmt"
)

var ErrGetOsName = errors.New("ErrGetOsName")

type errGetOsName struct {
	inner error
	msg   string
}

func (e *errGetOsName) Error() string {
	return fmt.Sprintf("%s Caused by: %s", e.msg, e.inner.Error())
}

func (e *errGetOsName) Unwrap() error {
	return e.inner
}

func (_ *errGetOsName) Is(target error) bool {
	return target == ErrGetOsName
}

func wrapErr(e error, msg string) error {
	return &errGetOsName{
		inner: e,
		msg:   msg,
	}
}

func OsName() (string, error) {
	s, e := osname()
	if e != nil {
		return "", wrapErr(e, "ErrGetOsName")
	}
	return s, nil
}
