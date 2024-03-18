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

type GetSchemaVariantFilter struct {
	ID       uuid.NullUUID
	SchemaID uuid.NullUUID
}

func (f *GetSchemaVariantFilter) toDataset() *goqu.SelectDataset {
	query := goqu.From(tbl_SchemaVariant)

	if f == nil {
		return query.Where(sqlpkg.UnrealCondition)
	}

	if f.ID.Valid {
		query = query.Where(col_SchemaVariant_ID.Eq(f.ID.UUID))
	}
	if f.SchemaID.Valid {
		query = query.Where(col_SchemaVariant_RefSchemaID.Eq(f.SchemaID.UUID))
	}

	return query
}

func (r *repositoryImpl) GetSchemaVariant(ctx context.Context, dbtx sqlpkg.DBTx, filter *GetSchemaVariantFilter) (*SchemaVariant, error) {
	const fn = "repositoryImpl.GetSchemaVariant"

	query, _, err := filter.toDataset().ToSQL()
	if err != nil {
		panic(err)
	}

	var ret SchemaVariant
	err = dbtx.GetContext(ctx, &ret, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Join(storage.ErrEntityNotFound, err)
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}

func (r *repositoryImpl) CreateSchemaVariant(ctx context.Context, dbtx sqlpkg.DBTx, record SchemaVariant) (*SchemaVariant, error) {
	const fn = "repositoryImpl.CreateSchemaVariant"

	timeNow := timepkg.TimeNow()
	record.CreatedAt = timeNow
	record.UpdatedAt = timeNow

	query, _, err := goqu.Insert(tbl_SchemaVariant).Rows(record).Returning(tbl_SchemaVariant.All()).ToSQL()
	if err != nil {
		panic(err)
	}

	var ret SchemaVariant
	err = dbtx.GetContext(ctx, &ret, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}
