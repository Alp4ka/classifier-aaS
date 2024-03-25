//	@Summary		Обновить состояние счета
//	@Description	Обновить состояние счета  и уведомить клиента.
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	HandlerResp		"Состояние счета обновлено"
//	@Failure		400			{object}	HandlerResp		"Ошибка валидации запроса"
//	@Failure		500			{object}	HandlerResp	"Внутренняя ошибка сервиса"
//	@Router			/api/schema/{id}  [get]

package http

import (
	"errors"
	"github.com/Alp4ka/classifier-aaS/internal/schemacomponent"
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
