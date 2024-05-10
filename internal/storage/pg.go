package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"runtime/debug"
)

var ErrTxAlreadyExists = errors.New("context already contains tx")
var UnrealCondition = goqu.I("1=0")

type postgresStorage struct {
	db *sqlx.DB
}

func NewPostgresStorage(db *sqlx.DB) Storage {
	return &postgresStorage{
		db: db,
	}
}

var _ Storage = (*postgresStorage)(nil)

func (s *postgresStorage) WithTransaction(ctx context.Context, function func(context.Context) error) (err error) {
	const fn = "storage.WithTransaction"

	var (
		tx    *sqlx.Tx = nil
		ctxTx          = txFromContext(ctx)
		hasTx          = ctxTx != nil
	)

	if hasTx {
		tx = ctxTx
	} else {
		tx, err = s.db.Beginx()
		if err != nil {
			return NewError(err).WithFn(fn)
		}
	}

	defer func() {
		if recovered := recover(); recovered != nil {
			recoveredErr, ok := recovered.(error)
			if !ok {
				recoveredErr = fmt.Errorf("panic val: %v", recovered)
			} else {
				recoveredErr = fmt.Errorf("panic error: %w", recoveredErr)
			}

			stacktrace := string(debug.Stack())
			err = fmt.Errorf("%w; stacktrace: %s", recoveredErr, stacktrace)
		}

		var storageError *Error
		if err != nil && !errors.As(err, &storageError) {
			err = NewError(err).WithFn(fn)
		}

		rollbackErr := tx.Rollback()
		if rollbackErr == nil || errors.Is(rollbackErr, sql.ErrTxDone) {
			return
		}

		err = errors.Join(err, rollbackErr)
	}()

	ctx = TxWithContext(ctx, tx)
	err = function(ctx)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *postgresStorage) DB() *sqlx.DB {
	return s.db
}

func (s *postgresStorage) DBTx(ctx context.Context) (dbtx DBTx, isTx bool) {
	tx := txFromContext(ctx)
	if tx != nil {
		return tx, true
	}

	return s.DB(), false
}

func (s *postgresStorage) GoquDBTx(ctx context.Context) (goquDBTx GoquDBTx, isTx bool) {
	tx := txFromContext(ctx)
	if tx != nil {
		return goqu.NewTx(s.Dialect(), tx), true
	}

	return goqu.New(s.Dialect(), s.DB()), false
}

func (s *postgresStorage) Dialect() string {
	return "postgres"
}
