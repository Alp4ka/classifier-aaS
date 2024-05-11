package repository

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

var tbl_Event = goqu.T("event")

type Event struct {
	ID           int64     `json:"id" db:"id" goqu:"skipinsert,skipupdate"`
	SessionID    uuid.UUID `json:"sessionID" db:"session_id" goqu:"skipupdate"`
	Req          string    `json:"req" db:"req" goqu:"skipupdate"`
	Resp         string    `json:"resp" db:"resp" goqu:"skipupdate"`
	SchemaNodeID uuid.UUID `json:"schemaNodeID" db:"schema_node_id" goqu:"skipupdate"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at" goqu:"skipupdate"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}
