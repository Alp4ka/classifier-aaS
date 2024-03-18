package api

import (
	"errors"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
	"github.com/Alp4ka/classifier-aaS/internal/schemacomponent"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
)

type hUpdateSchemaReq struct {
	ID          uuid.UUID           `json:"id"`
	Description *schema.Description `json:"description"`
	Gateway     null.String         `json:"gateway"`
}

func (s *HTTPServer) hUpdateSchema(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := new(hUpdateSchemaReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	schemaModel, err := s.schemaService.UpdateSchema(ctx,
		&schemacomponent.UpdateSchemaParams{
			ID:          req.ID,
			Gateway:     req.Gateway,
			Description: req.Description,
		},
	)
	if err != nil {
		if errors.Is(err, schemacomponent.ErrInvalidDescription) {
			return err
		}
		mlogger.L(ctx).Error("Error while creating schema", field.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.
		Status(fiber.StatusOK).
		JSON(
			HandlerResp{
				Success: true,
				Data:    schemaModel,
			},
		)
}

var _ fiber.Handler = (*HTTPServer)(nil).hUpdateSchema
