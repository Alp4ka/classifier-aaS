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
	"github.com/guregu/null/v5"
)

type GetSchemaFilter struct {
	ID      uuid.NullUUID
	Gateway null.String
}

func (f *GetSchemaFilter) toDataset() *goqu.SelectDataset {
	query := goqu.Select(tbl_Schema)

	if f == nil {
		return query.Where(sqlpkg.UnrealCondition)
	}

	if f.ID.Valid {
		query = query.Where(col_Schema_ID.Eq(f.ID.UUID))
	}
	if f.Gateway.Valid {
		query = query.Where(col_Schema_Gateway.Eq(f.Gateway.String))
	}

	return query
}

func (r *repositoryImpl) GetSchema(ctx context.Context, dbtx sqlpkg.DBTx, filter *GetSchemaFilter) (*Schema, error) {
	const fn = "repositoryImpl.GetSchema"

	query, _, err := filter.toDataset().ToSQL()
	if err != nil {
		panic(err)
	}

	var ret Schema
	err = dbtx.GetContext(ctx, &ret, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Join(storage.ErrEntityNotFound, err)
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}

func (r *repositoryImpl) CreateSchema(ctx context.Context, dbtx sqlpkg.DBTx, schema Schema) error {
	const fn = "repositoryImpl.CreateSchema"

	timeNow := timepkg.TimeNow()
	schema.CreatedAt = timeNow
	schema.UpdatedAt = timeNow

	query, _, err := goqu.Insert(tbl_Schema).Rows(schema).ToSQL()
	if err != nil {
		panic(err)
	}

	_, err = dbtx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (r *repositoryImpl) UpdateSchema(ctx context.Context, tx sqlpkg.Tx, schema Schema) error {
	const (
		fn                   = "repositoryImpl.UpdateSchema"
		expectedAffectedRows = 1
	)

	timeNow := timepkg.TimeNow()
	schema.UpdatedAt = timeNow

	query, _, err := goqu.Update(tbl_Schema).Set(schema).ToSQL()
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
