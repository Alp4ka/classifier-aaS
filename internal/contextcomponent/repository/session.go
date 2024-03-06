package repository

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID        uuid.UUID    `db:"id" goqu:"skipinsert,skipupdate"`
	ContextID uuid.UUID    `db:"context_id"`
	CreatedAt time.Time    `db:"created_at" goqu:"skipupdate"`
	UpdatedAt time.Time    `db:"updated_at"`
	State     SessionState `db:"state"`
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
