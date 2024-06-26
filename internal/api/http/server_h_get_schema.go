package http

import (
	"errors"

	schemacomponent "github.com/Alp4ka/classifier-aaS/internal/components/schema"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Server) hGetActualSchema(c *fiber.Ctx) error {
	ctx := c.UserContext()

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	schema, err := s.schemaService.GetSchema(
		c.Context(),
		&schemacomponent.GetSchemaFilter{
			ID: uuid.NullUUID{UUID: id, Valid: true},
		},
	)
	if err != nil {
		if errors.Is(err, schemacomponent.ErrSchemaNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}

		mlogger.L(ctx).Error("Error while getting schema", field.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.
		Status(fiber.StatusOK).
		JSON(
			HandlerResp{
				Success: true,
				Data:    schema,
			},
		)
}

var _ fiber.Handler = (*Server)(nil).hGetActualSchema
