package schemacomponent

import (
	"context"
	"errors"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/schema"
	"github.com/Alp4ka/classifier-aaS/internal/schemacomponent/repository"
	"github.com/Alp4ka/classifier-aaS/internal/storage"
	sqlpkg "github.com/Alp4ka/classifier-aaS/pkg/sql"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
)

type Config struct {
	Repository repository.Repository
}

type serviceImpl struct {
	repo repository.Repository
}

func NewService(cfg Config) Service {
	return &serviceImpl{
		repo: cfg.Repository,
	}
}

type GetSchemaFilter struct {
	ID     uuid.NullUUID
	Latest null.Bool
}

func (s *serviceImpl) GetSchema(ctx context.Context, filter *GetSchemaFilter) (*schema.Schema, error) {
	var ret *schema.Schema
	err := s.repo.WithTransaction(ctx,
		func(ctx context.Context, tx sqlpkg.Tx) error {
			schemaRecord, err := s.repo.GetSchema(ctx, tx,
				&repository.GetSchemaFilter{
					ID:     filter.ID,
					Latest: filter.Latest,
				},
			)
			if err != nil {
				if errors.Is(err, storage.ErrEntityNotFound) {
					return ErrSchemaNotFound
				}
				return err
			}

			ret = &schema.Schema{
				ID:        schemaRecord.ID,
				CreatedAt: schemaRecord.CreatedAt,
				UpdatedAt: schemaRecord.UpdatedAt,
			}
			if !schemaRecord.ActualVariantID.Valid {
				return nil
			}

			schemaVariantRecord, err := s.repo.GetSchemaVariant(ctx, tx,
				&repository.GetSchemaVariantFilter{
					ID: schemaRecord.ActualVariantID,
				},
			)
			if err != nil {
				return err
			}
			ret.ActualVariant = &schema.Variant{
				ID:          schemaVariantRecord.ID,
				Description: schemaVariantRecord.Description,
				CreatedAt:   schemaVariantRecord.CreatedAt,
				UpdatedAt:   schemaVariantRecord.UpdatedAt,
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

type CreateSchemaParams struct {
}

func (s *serviceImpl) CreateSchema(ctx context.Context, params *CreateSchemaParams) (*schema.Schema, error) {
	var ret *schema.Schema
	err := s.repo.WithTransaction(ctx,
		func(ctx context.Context, tx sqlpkg.Tx) error {
			schemaVariantRecord, err := s.repo.CreateSchemaVariant(ctx, tx,
				repository.SchemaVariant{
					ID:          uuid.New(),
					Description: schema.Description{},
					Editable:    true,
				},
			)
			if err != nil {
				return err
			}

			schemaRecord, err := s.repo.CreateSchema(ctx, tx,
				repository.Schema{
					ID:              uuid.New(),
					ActualVariantID: uuid.NullUUID{UUID: schemaVariantRecord.ID, Valid: true},
				},
			)
			if err != nil {
				return err
			}

			ret = &schema.Schema{
				ID:        schemaRecord.ID,
				CreatedAt: schemaRecord.CreatedAt,
				UpdatedAt: schemaRecord.UpdatedAt,
				ActualVariant: &schema.Variant{
					ID:          schemaVariantRecord.ID,
					Description: schemaVariantRecord.Description,
					CreatedAt:   schemaVariantRecord.CreatedAt,
					UpdatedAt:   schemaVariantRecord.UpdatedAt,
				},
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

type UpdateSchemaParams struct {
	ID          uuid.UUID           `json:"id"`
	Description *schema.Description `json:"description"`
}

func (s *serviceImpl) UpdateSchema(ctx context.Context, params *UpdateSchemaParams) (*schema.Schema, error) {
	var ret *schema.Schema

	if params.Description != nil {
		_, err := params.Description.MapAndValidate()
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrInvalidDescription, err.Error())
		}
	}

	err := s.repo.WithTransaction(ctx,
		func(ctx context.Context, tx sqlpkg.Tx) error {
			schemaRecord, err := s.repo.GetSchema(ctx, tx,
				&repository.GetSchemaFilter{
					ID: uuid.NullUUID{UUID: params.ID, Valid: true},
				},
			)
			if err != nil {
				if errors.Is(err, storage.ErrEntityNotFound) {
					return ErrSchemaNotFound
				}
				return err
			}

			updateSchemaVariantParams := repository.SchemaVariant{
				ID:       uuid.New(),
				Editable: true,
			}
			if params.Description != nil {
				updateSchemaVariantParams.Description = *(params.Description)
			}
			schemaVariantRecord, err := s.repo.CreateSchemaVariant(ctx, tx, updateSchemaVariantParams)
			if err != nil {
				return err
			}

			updateSchemaParams := repository.Schema{
				ID:              schemaRecord.ID,
				ActualVariantID: uuid.NullUUID{UUID: schemaVariantRecord.ID, Valid: true},
			}
			schemaRecord, err = s.repo.UpdateSchema(ctx, tx, updateSchemaParams)
			if err != nil {
				return err
			}

			ret = &schema.Schema{
				ID:        schemaRecord.ID,
				CreatedAt: schemaRecord.CreatedAt,
				UpdatedAt: schemaRecord.UpdatedAt,
				ActualVariant: &schema.Variant{
					ID:          schemaVariantRecord.ID,
					Description: schemaVariantRecord.Description,
					CreatedAt:   schemaVariantRecord.CreatedAt,
					UpdatedAt:   schemaVariantRecord.UpdatedAt,
				},
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *serviceImpl) GetSchemaVariant(ctx context.Context, id uuid.UUID) (*schema.Variant, error) {
	schemaVariantRecord, err := s.repo.GetSchemaVariant(ctx, s.repo.DB(),
		&repository.GetSchemaVariantFilter{
			ID: uuid.NullUUID{UUID: id, Valid: true},
		},
	)
	if err != nil {
		if errors.Is(err, storage.ErrEntityNotFound) {
			return nil, ErrVariantNotFound
		}
		return nil, err
	}

	return &schema.Variant{
		ID:          schemaVariantRecord.ID,
		Description: schemaVariantRecord.Description,
		CreatedAt:   schemaVariantRecord.CreatedAt,
		UpdatedAt:   schemaVariantRecord.UpdatedAt,
	}, nil
}

var _ Service = (*serviceImpl)(nil)
