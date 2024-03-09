package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/storage"
	sqlpkg "github.com/Alp4ka/classifier-aaS/pkg/sql"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type GetContextFilter struct {
	SessionID uuid.UUID
}

func (f *GetContextFilter) toDataset() *goqu.SelectDataset {
	query := goqu.Select(tbl_Context)

	if f == nil {
		query = query.Where(sqlpkg.UnrealCondition)
	} else {
		query = query.Where(col_Context_SessionID.Eq(f.SessionID))
	}

	return query
}

func (r *repositoryImpl) GetContext(ctx context.Context, dbtx sqlpkg.DBTx, filter *GetContextFilter) (*Context, error) {
	const fn = "repositoryImpl.GetContext"

	query, _, err := filter.toDataset().ToSQL()
	if err != nil {
		panic(err)
	}

	var ret Context
	err = dbtx.GetContext(ctx, &ret, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Join(storage.ErrEntityNotFound, sql.ErrNoRows)
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}

func (r *repositoryImpl) CreateContext(ctx context.Context, dbtx sqlpkg.DBTx, context Context) error {
	const fn = "repositoryImpl.CreateContext"

	timeNow := timepkg.TimeNow()
	context.CreatedAt = timeNow
	context.UpdatedAt = timeNow

	query, _, err := goqu.Insert(tbl_Context).Rows(context).ToSQL()
	if err != nil {
		panic(err)
	}

	_, err = dbtx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (r *repositoryImpl) UpdateContext(ctx context.Context, tx sqlpkg.Tx, context Context) error {
	const (
		fn                   = "repositoryImpl.UpdateContext"
		expectedAffectedRows = 1
	)

	timeNow := timepkg.TimeNow()
	context.UpdatedAt = timeNow

	query, _, err := goqu.Update(tbl_Context).Set(context).ToSQL()
	if err != nil {
		panic(err)
	}

	res, err := tx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	affected, err := res.RowsAffected()
	if err != nil || affected != expectedAffectedRows {
		return fmt.Errorf("%s: affected %d, expected %d; %w", fn, affected, expectedAffectedRows, err)
	}

	return nil
}
