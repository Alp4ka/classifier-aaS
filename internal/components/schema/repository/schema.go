package repository

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

var tbl_Schema = goqu.T("schema")

type Schema struct {
	ID              uuid.UUID     `db:"id" goqu:"skipupdate"`
	ActualVariantID uuid.NullUUID `db:"actual_variant_id"`
	CreatedAt       time.Time     `db:"created_at" goqu:"skipupdate"`
	UpdatedAt       time.Time     `db:"updated_at"`
}
