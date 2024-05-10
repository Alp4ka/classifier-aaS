package context

import "errors"

var (
	ErrNoActualVariant     = errors.New("no schema actual variant")
	ErrSessionDoesNotExist = errors.New("session does not exist")
)
