package repository

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Context struct {
	ID              uuid.UUID    `db:"id" goqu:"skipinsert,skipupdate"`
	SessionID       uuid.UUID    `db:"session_id"`
	SchemaVariantID uuid.UUID    `db:"schema_variant_id"`
	SchemaStepXPath string       `db:"schema_x_path"`
	State           ContextState `db:"state"`
	CreatedAt       time.Time    `db:"created_at" goqu:"skipupdate"`
	UpdatedAt       time.Time    `db:"updated_at"`
}

type ContextState string

func (cs ContextState) Value() (driver.Value, error) {
	return string(cs), nil
}

func (cs *ContextState) Scan(value interface{}) error {
	if value == nil {
		return errors.New("contextcomponent state nil value")
	}

	valueStr, ok := value.(string)
	if !ok {
		return errors.New("contextcomponent state was not a string")
	}

	if _, isAvailable := _availableContextStates[ContextState(valueStr)]; !isAvailable {
		return fmt.Errorf("unknown contextcomponent state '%s'", value.(string))
	}
	return nil
}

const (
	ContextStateActive ContextState = "active"
	ContextStateClosed ContextState = "closed"
	ContextStateError  ContextState = "error"
)

var _availableContextStates = map[ContextState]struct{}{
	ContextStateActive: {},
	ContextStateClosed: {},
	ContextStateError:  {},
}
