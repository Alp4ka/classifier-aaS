package repository

import (
	"context"
	"fmt"
	sqlpkg "github.com/Alp4ka/classifier-aaS/pkg/sql"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"
	"github.com/doug-martin/goqu/v9"
)

func (r *repositoryImpl) CreateEvent(ctx context.Context, dbtx sqlpkg.DBTx, event Event) (*Event, error) {
	const fn = "repositoryImpl.CreateEvent"

	event.CreatedAt = timepkg.TimeNow()
	event.UpdatedAt = timepkg.TimeNow()

	query, _, err := goqu.Insert(tbl_Event).Rows(event).Returning(tbl_Event.All()).ToSQL()
	if err != nil {
		panic(err)
	}

	var ret Event
	err = dbtx.GetContext(ctx, &ret, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}
