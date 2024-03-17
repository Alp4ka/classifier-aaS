package repository

import (
	"github.com/Alp4ka/classifier-aaS/internal/schema"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"time"
)

var (
	tbl_SchemaVariant = goqu.T("schema_variant")

	col_SchemaVariant_ID       = tbl_Schema.Col("id")
	col_SchemaVariant_SchemaID = tbl_Schema.Col("schema_id")
)

type SchemaVariant struct {
	ID          uuid.UUID          `json:"id" db:"id" goqu:"skipupdate"`
	RefSchemaID uuid.UUID          `json:"refSchemaID" db:"ref_schema_id" goqu:"skipupdate"`
	Description schema.Description `json:"description" db:"description"`
	CreatedAt   time.Time          `json:"createdAt" db:"created_at" goqu:"skipupdate"`
	UpdatedAt   time.Time          `json:"updatedAt" db:"updated_at"`
}
