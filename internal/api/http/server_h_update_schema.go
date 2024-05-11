package http

import (
	"errors"

	schemacomponent "github.com/Alp4ka/classifier-aaS/internal/components/schema"
	"github.com/Alp4ka/classifier-aaS/internal/components/schema/entities"
	"github.com/Alp4ka/mlogger"
	"github.com/Alp4ka/mlogger/field"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type hUpdateSchemaReq struct {
	ID          uuid.UUID             `json:"id"`
	Description *entities.Description `json:"description"`
}

func (s *Server) hUpdateSchema(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := new(hUpdateSchemaReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	schemaModel, err := s.schemaService.UpdateSchema(ctx,
		&schemacomponent.UpdateSchemaParams{
			ID:          req.ID,
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

var _ fiber.Handler = (*Server)(nil).hUpdateSchema
