package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SQLStorage interface {
	mustEmbedSQLStorageImpl()
	DB() *sqlx.DB
	WithTransaction(ctx context.Context, fn func(ctx context.Context, tx *sqlx.Tx) error) error
}

type SQLStorageImpl struct {
	Db *sqlx.DB
}

func (s *SQLStorageImpl) mustEmbedSQLStorageImpl() {}

func (s *SQLStorageImpl) DB() *sqlx.DB {
	return s.Db
}

func (s *SQLStorageImpl) WithTransaction(ctx context.Context, f func(context.Context, *sqlx.Tx) error) (err error) {
	const fn = "SQLStorageImpl.WithTransaction"
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s: %w", fn, err)
		}
	}()

	tx, err := s.DB().BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		recoverRet := recover()
		if recoverErr, ok := recoverRet.(error); ok {
			err = recoverErr
		}

		rollbackErr := tx.Rollback()
		if rollbackErr != nil && !errors.Is(rollbackErr, sql.ErrTxDone) {
			err = errors.Join(err, rollbackErr)
		}
	}()

	err = f(ctx, tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

var _ SQLStorage = (*SQLStorageImpl)(nil)
