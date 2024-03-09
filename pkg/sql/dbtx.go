package sql

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DBTx interface {
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

type Tx interface {
	DBTx
	Commit() error
	Rollback() error
}

var (
	_ DBTx = (*sqlx.DB)(nil)
	_ DBTx = (*sqlx.Tx)(nil)
	_ Tx   = (*sqlx.Tx)(nil)
)
