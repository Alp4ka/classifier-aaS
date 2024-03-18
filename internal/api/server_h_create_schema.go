package api

import (
	"github.com/Alp4ka/classifier-aaS/internal/schemacomponent"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/gofiber/fiber/v2"
)

// TODO: Validation.
type hCreateSchemaReq struct {
	Gateway string `json:"gateway" validate:"required"`
}

func (s *HTTPServer) hCreateSchema(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := new(hCreateSchemaReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	schema, err := s.schemaService.CreateSchema(c.Context(),
		&schemacomponent.CreateSchemaParams{
			Gateway: req.Gateway,
		},
	)
	if err != nil {
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

var _ fiber.Handler = (*HTTPServer)(nil).hCreateSchema