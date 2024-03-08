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

	col_Session_ID = tbl_Session.Col("id")
)

type Session struct {
	ID         uuid.UUID    `db:"id" goqu:"skipinsert,skipupdate"`
	CreatedAt  time.Time    `db:"created_at" goqu:"skipupdate"`
	ValidUntil time.Time    `db:"valid_until" goqu:"skipupdate"`
	UpdatedAt  time.Time    `db:"updated_at"`
	State      SessionState `db:"state"`
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
	SessionStateActive SessionState = "active"
	SessionStateClosed SessionState = "closed"
	SessionStateError  SessionState = "error"
)

var _availableSessionStates = map[SessionState]struct{}{
	SessionStateActive: {},
	SessionStateClosed: {},
	SessionStateError:  {},
}
