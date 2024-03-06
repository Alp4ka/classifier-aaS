package sql

import "errors"

var (
	ErrNoTxInCtx = errors.New("no transaction in contextcomponent")
)
