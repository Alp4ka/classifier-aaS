package sql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type dbtxer interface {
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type helper interface {
	mustEmbedDbtxWrapper()
	IsTx() bool
}

type DBTX interface {
	dbtxer
	helper
}

type dbtxWrapper struct {
	dbtxer
	isTx bool
}

func (d *dbtxWrapper) IsTx() bool {
	return d.isTx
}

func (d *dbtxWrapper) mustEmbedDbtxWrapper() {}

func WrapDB(db *sqlx.DB) DBTX {
	return &dbtxWrapper{dbtxer: db, isTx: false}
}

func WrapTx(tx *sqlx.Tx) DBTX {
	return &dbtxWrapper{dbtxer: tx, isTx: true}
}

var (
	_ dbtxer = (*sqlx.DB)(nil)
	_ dbtxer = (*sqlx.Tx)(nil)
)
