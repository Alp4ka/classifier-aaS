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
	"time"
)

type GetSessionFilter struct {
	ID          uuid.UUID
	CurrentTime time.Time
	States      []SessionState
}

func (f *GetSessionFilter) toDataset() *goqu.SelectDataset {
	if f == nil {
		return nil
	}

	query := goqu.
		From(tbl_Session).
		Where(col_Session_ID.Eq(f.ID)).
		Where(col_Session_ValidUntil.Gte(f.CurrentTime))
	if f.States != nil {
		query = query.Where(col_Session_State.In(f.States))
	}

	return query
}

func (r *repositoryImpl) GetSession(ctx context.Context, dbtx sqlpkg.DBTx, filter *GetSessionFilter) (*Session, error) {
	const fn = "repositoryImpl.GetSession"

	query, _, err := filter.toDataset().ToSQL()
	if err != nil {
		panic(err)
	}

	var session Session
	err = dbtx.GetContext(ctx, &session, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Join(storage.ErrEntityNotFound, err)
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &session, nil
}

func (r *repositoryImpl) CreateSession(ctx context.Context, dbtx sqlpkg.DBTx, session Session) (*Session, error) {
	const fn = "repositoryImpl.CreateSession"

	timeNow := timepkg.TimeNow()
	session.CreatedAt = timeNow
	session.UpdatedAt = timeNow

	query, _, err := goqu.Insert(tbl_Session).Rows(session).Returning(tbl_Session.All()).ToSQL()
	if err != nil {
		panic(err)
	}

	var ret Session
	err = dbtx.GetContext(ctx, &ret, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}

func (r *repositoryImpl) UpdateSession(ctx context.Context, tx sqlpkg.Tx, session Session) (*Session, error) {
	const fn = "repositoryImpl.UpdateSession"

	session.UpdatedAt = timepkg.TimeNow()

	query, _, err := goqu.Update(tbl_Session).Set(session).Returning(tbl_Session.All()).ToSQL()
	if err != nil {
		panic(err)
	}

	var ret Session
	err = tx.GetContext(ctx, &ret, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrEntityNotFound
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}
