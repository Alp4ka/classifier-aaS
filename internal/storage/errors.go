package storage

import (
	"errors"
	"fmt"
)

var ErrEntityNotFound = errors.New("entity not found")

type Error struct {
	fn  string
	err error
}

func NewError(err error) *Error {
	return &Error{err: err}
}

func NewQueryError(err error, query string) *Error {
	return &Error{err: fmt.Errorf("err=%w; query=%s", err, query)}
}

func (e *Error) WithFn(fn string) *Error {
	e.fn = fn
	return e
}

func (e *Error) Error() string {
	const emptyErrStr = "empty storage error"

	ret := emptyErrStr
	if e.err != nil {
		ret = fmt.Sprintf("%s storage error", e.err.Error())
	}

	if e.fn != "" {
		ret = fmt.Sprintf("%s: ", e.fn) + ret
	}

	return ret
}

func (e *Error) Unwrap() error {
	return e.err
}
