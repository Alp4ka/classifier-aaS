package repository

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"time"
)

var (
	tbl_Schema = goqu.T("schema")

	col_Schema_ID      = tbl_Schema.Col("id")
	col_Schema_Gateway = tbl_Schema.Col("gateway")
)

type Schema struct {
	ID              uuid.UUID     `db:"id" goqu:"skipupdate"`
	Gateway         string        `db:"gateway"`
	ActualVariantID uuid.NullUUID `db:"actual_variant_id"`
	CreatedAt       time.Time     `db:"created_at" goqu:"skipupdate"`
	UpdatedAt       time.Time     `db:"updated_at"`
}
