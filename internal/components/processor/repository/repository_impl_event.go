package repository

import (
	"context"
	"fmt"
	timepkg "github.com/Alp4ka/classifier-aaS/pkg/time"
	"github.com/doug-martin/goqu/v9"
)

func (r *repositoryImpl) CreateEvent(ctx context.Context, event Event) (*Event, error) {
	const fn = "repositoryImpl.CreateEvent"

	event.CreatedAt = timepkg.Now()
	event.UpdatedAt = timepkg.Now()

	query, _, err := goqu.Insert(tbl_Event).Rows(event).Returning(tbl_Event.All()).ToSQL()
	if err != nil {
		panic(err)
	}

	var ret Event
	executor, _ := r.DBTx(ctx)
	err = executor.GetContext(ctx, &ret, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &ret, nil
}
