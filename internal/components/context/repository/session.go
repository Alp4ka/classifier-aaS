package repository

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
)

var tbl_Session = goqu.T("session")

type Session struct {
	ID         uuid.UUID    `json:"id" db:"id" goqu:"skipupdate"`
	State      SessionState `json:"state" db:"state"`
	Agent      string       `json:"agent" db:"agent" goqu:"skipupdate"`
	Gateway    string       `json:"gateway" db:"gateway" goqu:"skipupdate"`
	ValidUntil time.Time    `json:"validUntil" db:"valid_until" goqu:"skipupdate"`
	ClosedAt   null.Time    `json:"closedAt" db:"closed_at"`

	SchemaVariantID uuid.UUID `json:"schemaVariantID" db:"schema_variant_id" goqu:"skipupdate"`
	SchemaNodeID    uuid.UUID `json:"schemaNodeID" db:"schema_node_id"`

	CreatedAt time.Time `json:"createdAt" db:"created_at" goqu:"skipupdate"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type SessionState string

func (cs SessionState) Value() (driver.Value, error) {
	return string(cs), nil
}

func (cs *SessionState) Scan(value interface{}) error {
	if value == nil {
		return errors.New("session state nil value")
	}

	valueStr, ok := value.(string)
	if !ok {
		return errors.New("session state was not a string")
	}

	if _, isAvailable := _availableSessionStates[SessionState(valueStr)]; !isAvailable {
		return fmt.Errorf("unknown session state '%s'", value.(string))
	}

	*cs = SessionState(valueStr)
	return nil
}

const (
	SessionStateActive        SessionState = "active"
	SessionStateClosedAgent   SessionState = "closed_by_agent"
	SessionStateClosedGateway SessionState = "closed_by_gateway"
	SessionStateFinished      SessionState = "finished"
	SessionStateClosedRotten  SessionState = "closed_by_rot"
	SessionStateError         SessionState = "error"
)

var _availableSessionStates = map[SessionState]struct{}{
	SessionStateActive:        {},
	SessionStateClosedAgent:   {},
	SessionStateClosedGateway: {},
	SessionStateFinished:      {},
	SessionStateClosedRotten:  {},
	SessionStateError:         {},
}
