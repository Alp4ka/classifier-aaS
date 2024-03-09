package repository

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"time"
)

var tbl_Event = goqu.T("event")

type Event struct {
	ID        int64     `json:"id" db:"id" goqu:"skipinsert,skipupdate"`
	ContextID uuid.UUID `json:"context_id" db:"context_id" goqu:"skipupdate"`
	Req       *string   `json:"req" db:"req" goqu:"skipupdate"`
	Resp      *string   `json:"resp" db:"resp" goqu:"skipupdate"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" goqu:"skipupdate"`
}
