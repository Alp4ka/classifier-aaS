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
	query := goqu.Select(tbl_SchemaVariant)

	if f == nil {
		return query.Where(sqlpkg.UnrealCondition)
	}

	if f.ID.Valid {
		query = query.Where(col_SchemaVariant_ID.Eq(f.ID.UUID))
	}
	if f.SchemaID.Valid {
		query = query.Where(col_SchemaVariant_SchemaID.Eq(f.SchemaID.UUID))
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

func (r *repositoryImpl) CreateSchemaVariant(ctx context.Context, dbtx sqlpkg.DBTx, record SchemaVariant) error {
	const fn = "repositoryImpl.CreateSchemaVariant"

	timeNow := timepkg.TimeNow()
	record.CreatedAt = timeNow
	record.UpdatedAt = timeNow

	query, _, err := goqu.Insert(tbl_SchemaVariant).Rows(record).ToSQL()
	if err != nil {
		panic(err)
	}

	_, err = dbtx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (r *repositoryImpl) UpdateSchemaVariant(ctx context.Context, tx sqlpkg.Tx, record SchemaVariant) error {
	const (
		fn                   = "repositoryImpl.UpdateSchemaVariant"
		expectedAffectedRows = 1
	)

	timeNow := timepkg.TimeNow()
	record.UpdatedAt = timeNow

	query, _, err := goqu.Update(tbl_SchemaVariant).Set(record).ToSQL()
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
