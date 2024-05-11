package repository

import (
	"time"

	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

var tbl_SchemaVariant = goqu.T("schema_variant")

type SchemaVariant struct {
	ID          uuid.UUID            `db:"id" goqu:"skipupdate"`
	Description entities.Description `db:"description"`
	Editable    bool                 `db:"editable"`
	CreatedAt   time.Time            `db:"created_at" goqu:"skipupdate"`
	UpdatedAt   time.Time            `db:"updated_at"`
}
