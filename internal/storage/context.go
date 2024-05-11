package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
)

var _txKey struct{}

func txWithContext(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, _txKey, tx)
}

func txFromContext(ctx context.Context) *sqlx.Tx {
	tx, _ := ctx.Value(_txKey).(*sqlx.Tx)
	return tx
}
