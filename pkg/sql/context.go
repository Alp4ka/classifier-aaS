package sql

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type ctxTxKey struct{}

var _ctxTxKeyVal = ctxTxKey{}

func TxFromCtx(ctx context.Context) (*sqlx.Tx, error) {
	tx, ok := ctx.Value(_ctxTxKeyVal).(*sqlx.Tx)
	if !ok {
		return nil, ErrNoTxInCtx
	}
	return tx, nil
}

func CtxWithTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, _ctxTxKeyVal, tx)
}
