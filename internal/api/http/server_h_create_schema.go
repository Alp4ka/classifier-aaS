package http

import (
	"errors"
	schemacomponent "github.com/Alp4ka/classifier-aaS/internal/components/schema"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/gofiber/fiber/v2"
)

type hCreateSchemaReq struct {
}

func (s *Server) hCreateSchema(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := new(hCreateSchemaReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	schema, err := s.schemaService.CreateSchema(c.Context(), &schemacomponent.CreateSchemaParams{})
	if err != nil {
		if errors.Is(err, schemacomponent.ErrOnlySingleSchemaAvailable) {
			return c.
				Status(fiber.StatusConflict).
				JSON(
					HandlerResp{
						Success: false,
						Message: "single schema available!",
					},
				)
		}
		mlogger.L(ctx).Error("Error while creating schema", field.Error(err))
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

var _ fiber.Handler = (*Server)(nil).hCreateSchema
