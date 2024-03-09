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
	tbl_Context = goqu.T("context")

	col_Context_SessionID = tbl_Session.Col("id")
)

type Context struct {
	ID              uuid.UUID              `json:"id" db:"id" goqu:"skipupdate"`
	SessionID       uuid.UUID              `json:"sessionID" db:"session_id" goqu:"skipupdate"`
	SchemaVariantID uuid.UUID              `json:"schemaVariantID" db:"schema_variant_id" goqu:"skipupdate"`
	SchemaStepXPath string                 `json:"schemaStepXPath" db:"schema_x_path"`
	Metadata        map[string]interface{} `json:"metadata" db:"metadata"`
	State           ContextState           `json:"state" db:"state"`
	CreatedAt       time.Time              `json:"createdAt" db:"created_at" goqu:"skipupdate"`
	UpdatedAt       time.Time              `json:"updatedAt" db:"updated_at"`
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
