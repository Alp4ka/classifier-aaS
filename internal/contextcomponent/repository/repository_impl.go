package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Alp4ka/classifier-aaS/internal/storage"
	sqlpkg "github.com/Alp4ka/classifier-aaS/pkg/sql"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type repositoryImpl struct {
	*storage.SQLStorageImpl
}

type GetSessionFilter struct {
	ID uuid.UUID
}

func (f *GetSessionFilter) toDataset() *goqu.SelectDataset {
	var id uuid.UUID
	if f == nil || uuid.Validate(f.ID.String()) != nil {
		id = uuid.New()
	} else {
		id = f.ID
	}

	return goqu.Select(tbl_Session).Where(col_Session_ID.Eq(id))
}

func (r *repositoryImpl) GetSession(ctx context.Context, dbtx sqlpkg.DBTX, filter *GetSessionFilter) (*Session, error) {
	const fn = "repositoryImpl.GetSession"

	query, _, err := filter.toDataset().ToSQL()
	if err != nil {
		panic(err)
	}

	var session Session
	err = dbtx.GetContext(ctx, &session, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrEntityNotFound
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &session, nil
}

func (r *repositoryImpl) CreateSession(ctx context.Context, dbtx sqlpkg.DBTX, session Session) error {
	const fn = "repositoryImpl.CreateSession"

	timeNow := timepkg.TimeNow()
	session.CreatedAt = timeNow
	session.UpdatedAt = timeNow

	query, _, err := goqu.Insert(tbl_Session).Rows(session).ToSQL()
	if err != nil {
		panic(err)
	}

	_, err = dbtx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

var _ Repository = (*repositoryImpl)(nil)
