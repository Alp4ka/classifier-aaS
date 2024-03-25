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
	"github.com/guregu/null/v5"
)

type GetSessionFilter struct {
	ID      uuid.NullUUID
	Gateway null.String
	Agent   null.String
	Active  null.Bool
}

func (f *GetSessionFilter) toDataset() *goqu.SelectDataset {
	query := goqu.From(tbl_Session)

	if f.ID.Valid {
		query = query.Where(tbl_Session.Col("id").Eq(f.ID.UUID))
	}

	if f.Gateway.Valid {
		query = query.Where(tbl_Session.Col("gateway").Eq(f.Gateway.String))
	}

	if f.Agent.Valid {
		query = query.Where(tbl_Session.Col("agent").Eq(f.Agent.String))
	}

	if f.Agent.Valid || f.Gateway.Valid {
		query.Order(tbl_Session.Col("created_at").Asc()).Limit(1)
	}

	if f.Active.Valid {
		if !f.Active.Bool {
			panic("unexpected filter active was false")
		}
		query = query.
			Where(tbl_Session.Col("valid_until").Gt(timepkg.Now())).
			Where(tbl_Session.Col("state").Eq(SessionStateActive)).
			Where(tbl_Session.Col("closed_at").IsNull())
	}

	return query
}

func (r *repositoryImpl) GetSession(ctx context.Context, dbtx sqlpkg.DBTx, filter *GetSessionFilter) (*Session, error) {
	const fn = "repositoryImpl.GetSession"

	query, _, err := filter.toDataset().ToSQL()
	if err != nil {
		panic(err)
	}

	var ret Session
	err = dbtx.GetContext(ctx, &ret, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Join(err, storage.ErrEntityNotFound)
		}
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}

func (r *repositoryImpl) CreateSession(ctx context.Context, dbtx sqlpkg.DBTx, session Session) (*Session, error) {
	const fn = "repositoryImpl.CreateSession"

	timeNow := timepkg.Now()
	session.CreatedAt = timeNow
	session.UpdatedAt = timeNow

	query, _, err := goqu.Insert(tbl_Session).Rows(session).OnConflict(goqu.DoNothing()).Returning(tbl_Session.All()).ToSQL()
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

	session.UpdatedAt = timepkg.Now()

	query, _, err := goqu.Update(tbl_Session).Set(session).Where(tbl_Session.Col("id").Eq(session.ID)).Returning(tbl_Session.All()).ToSQL()
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
