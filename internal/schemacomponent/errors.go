package schemacomponent

import "errors"

var (
	ErrSchemaNotFound     = errors.New("schema not found")
	ErrVariantNotFound    = errors.New("variant not found")
	ErrInvalidDescription = errors.New("invalid description")
)
