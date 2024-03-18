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
	query := goqu.From(tbl_Context)

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
			return nil, errors.Join(storage.ErrEntityNotFound, err)
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}

func (r *repositoryImpl) CreateContext(ctx context.Context, dbtx sqlpkg.DBTx, context Context) (*Context, error) {
	const fn = "repositoryImpl.CreateContext"

	timeNow := timepkg.TimeNow()
	context.CreatedAt = timeNow
	context.UpdatedAt = timeNow

	query, _, err := goqu.Insert(tbl_Context).Rows(context).Returning(tbl_Context.All()).ToSQL()
	if err != nil {
		panic(err)
	}

	var ret Context
	err = dbtx.GetContext(ctx, &ret, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}

func (r *repositoryImpl) UpdateContext(ctx context.Context, tx sqlpkg.Tx, context Context) (*Context, error) {
	const fn = "repositoryImpl.UpdateContext"

	timeNow := timepkg.TimeNow()
	context.UpdatedAt = timeNow

	query, _, err := goqu.Update(tbl_Context).Set(context).Returning(tbl_Context.All()).ToSQL()
	if err != nil {
		panic(err)
	}

	var ret Context
	err = tx.GetContext(ctx, &ret, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrEntityNotFound
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}
