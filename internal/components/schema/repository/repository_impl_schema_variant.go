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
)

type GetSchemaVariantFilter struct {
	ID uuid.NullUUID
}

func (f *GetSchemaVariantFilter) toDataset() *goqu.SelectDataset {
	query := goqu.From(tbl_SchemaVariant)

	if f == nil {
		return query.Where(storage.UnrealCondition)
	}

	if f.ID.Valid {
		query = query.Where(tbl_SchemaVariant.Col("id").Eq(f.ID.UUID))
	}

	return query
}

func (r *repositoryImpl) GetSchemaVariant(ctx context.Context, filter *GetSchemaVariantFilter) (*SchemaVariant, error) {
	const fn = "repositoryImpl.GetSchemaVariant"

	query, _, err := filter.toDataset().ToSQL()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	var ret SchemaVariant
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

func (r *repositoryImpl) CreateSchemaVariant(ctx context.Context, record SchemaVariant) (*SchemaVariant, error) {
	const fn = "repositoryImpl.CreateSchemaVariant"

	timeNow := timepkg.Now()
	record.CreatedAt = timeNow
	record.UpdatedAt = timeNow

	query, _, err := goqu.Insert(tbl_SchemaVariant).Rows(record).Returning(tbl_SchemaVariant.All()).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	var ret SchemaVariant
	executor, _ := r.DBTx(ctx)
	err = executor.GetContext(ctx, &ret, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}
