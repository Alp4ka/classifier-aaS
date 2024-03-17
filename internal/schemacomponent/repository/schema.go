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
	ID              uuid.UUID `json:"id" db:"id" goqu:"skipupdate"`
	Gateway         string    `json:"gateway" db:"gateway"`
	ActualVariantID uuid.UUID `json:"actualVariantID" db:"actual_variant_id"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at" goqu:"skipupdate"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
}
