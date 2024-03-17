package repository

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"time"
)

var (
	tbl_Session = goqu.T("session")

	col_Session_ID         = tbl_Session.Col("id")
	col_Session_ValidUntil = tbl_Session.Col("valid_until")
	col_Session_State      = tbl_Session.Col("state")
)

type Session struct {
	ID         uuid.UUID    `json:"id" db:"id" goqu:"skipupdate"`
	State      SessionState `json:"state" db:"state"`
	Agent      string       `json:"agent" db:"agent" goqu:"skipupdate"`
	Gateway    string       `json:"gateway" db:"gateway" goqu:"skipupdate"`
	ValidUntil time.Time    `json:"validUntil" db:"valid_until" goqu:"skipupdate"` // TODO(Gorkovets Roman): logic for deletion invalid.
	ClosedAt   time.Time    `json:"closedAt" db:"closed_at"`
	CreatedAt  time.Time    `json:"createdAt" db:"created_at" goqu:"skipupdate"`
	UpdatedAt  time.Time    `json:"updatedAt" db:"updated_at"`
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
	return nil
}

const (
	SessionStateActive       SessionState = "active"
	SessionStateClosedAgent  SessionState = "closed_by_agent"
	SessionStateClosedRotten SessionState = "closed_by_rot"
	SessionStateError        SessionState = "error"
)

var _availableSessionStates = map[SessionState]struct{}{
	SessionStateActive:       {},
	SessionStateClosedAgent:  {},
	SessionStateClosedRotten: {},
	SessionStateError:        {},
}
