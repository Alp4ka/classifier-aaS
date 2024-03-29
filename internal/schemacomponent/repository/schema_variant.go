package repository

import (
	"github.com/Alp4ka/classifier-aaS/internal/schema"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"time"
)

var tbl_SchemaVariant = goqu.T("schema_variant")

type SchemaVariant struct {
	ID          uuid.UUID          `db:"id" goqu:"skipupdate"`
	Description schema.Description `db:"description"`
	Editable    bool               `db:"editable"`
	CreatedAt   time.Time          `db:"created_at" goqu:"skipupdate"`
	UpdatedAt   time.Time          `db:"updated_at"`
}
