package schemacomponent

import "errors"

var (
	ErrSchemaNotFound     = errors.New("schema not found")
	ErrInvalidDescription = errors.New("invalid description")
)
