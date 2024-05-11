package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/storage"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
)

type GetSchemaFilter struct {
	ID     uuid.NullUUID
	Latest null.Bool
}

func (f *GetSchemaFilter) toDataset() *goqu.SelectDataset {
	query := goqu.From(tbl_Schema)

	if f == nil {
		return query.Where(storage.UnrealCondition)
	}

	if f.ID.Valid {
		query = query.Where(tbl_Schema.Col("id").Eq(f.ID.UUID))
	}

	if f.Latest.Valid {
		query = query.Order(tbl_Schema.Col("created_at").Desc()).Limit(1)
	}

	return query
}

func (r *repositoryImpl) GetSchema(ctx context.Context, filter *GetSchemaFilter) (*Schema, error) {
	const fn = "repositoryImpl.GetSchema"

	query, _, err := filter.toDataset().ToSQL()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	var ret Schema
	executor, _ := r.DBTx(ctx)
	err = executor.GetContext(ctx, &ret, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Join(storage.ErrEntityNotFound, err)
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}

func (r *repositoryImpl) CreateSchema(ctx context.Context, schema Schema) (*Schema, error) {
	const fn = "repositoryImpl.CreateSchema"

	timeNow := timepkg.Now()
	schema.CreatedAt = timeNow
	schema.UpdatedAt = timeNow

	query, _, err := goqu.Insert(tbl_Schema).Rows(schema).Returning(tbl_Schema.All()).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	var ret Schema
	executor, _ := r.DBTx(ctx)
	err = executor.GetContext(ctx, &ret, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}

func (r *repositoryImpl) UpdateSchema(ctx context.Context, schema Schema) (*Schema, error) {
	const fn = "repositoryImpl.UpdateSchema"

	timeNow := timepkg.Now()
	schema.UpdatedAt = timeNow

	query, _, err := goqu.Update(tbl_Schema).Set(schema).Returning(tbl_Schema.All()).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	var ret Schema
	executor, _ := r.DBTx(ctx)
	err = executor.GetContext(ctx, &ret, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrEntityNotFound
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}
