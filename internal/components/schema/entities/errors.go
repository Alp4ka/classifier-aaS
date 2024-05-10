package entities

import "errors"

var (
	ErrMissingField   = errors.New("missing field")
	ErrFieldWrongType = errors.New("field has wrong type")
)
