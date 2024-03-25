package contextcomponent

import "errors"

var (
	ErrNoActualVariant     = errors.New("no schema actual variant")
	ErrSessionExpired      = errors.New("session expired")
	ErrSessionDoesNotExist = errors.New("session does not exist")
)
