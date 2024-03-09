package repository

import (
	"context"
	"fmt"
	sqlpkg "github.com/Alp4ka/classifier-aaS/pkg/sql"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"
	"github.com/doug-martin/goqu/v9"
)

func (r *repositoryImpl) CreateEvent(ctx context.Context, dbtx sqlpkg.DBTx, event Event) error {
	const fn = "repositoryImpl.CreateEvent"

	event.CreatedAt = timepkg.TimeNow()

	query, _, err := goqu.Insert(tbl_Event).Rows(event).ToSQL()
	if err != nil {
		panic(err)
	}

	_, err = dbtx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
