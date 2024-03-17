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
		Select(tbl_Session).
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

func (r *repositoryImpl) CreateSession(ctx context.Context, dbtx sqlpkg.DBTx, session *Session) error {
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

func (r *repositoryImpl) UpdateSession(ctx context.Context, tx sqlpkg.Tx, session *Session) error {
	const (
		fn                   = "repositoryImpl.UpdateSession"
		expectedAffectedRows = 1
	)

	session.UpdatedAt = timepkg.TimeNow()

	query, _, err := goqu.Update(tbl_Session).Set(session).ToSQL()
	if err != nil {
		panic(err)
	}

	res, err := tx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	affected, err := res.RowsAffected()
	if err != nil || affected != expectedAffectedRows {
		return fmt.Errorf("%s: affected %d, expected %d; %w", fn, affected, expectedAffectedRows, err)
	}

	return nil
}
