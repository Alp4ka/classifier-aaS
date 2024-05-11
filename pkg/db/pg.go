package db

import (
	"fmt"

	"github.com/j2gg0s/otsql"
	"github.com/j2gg0s/otsql/hook/metric"
	"github.com/j2gg0s/otsql/hook/trace"
	"github.com/jmoiron/sqlx"
)

const DriverName = "pgx"

func NewPgDB(cfg Config) (*sqlx.DB, error) {
	metricHook, err := metric.New()
	if err != nil {
		return nil, fmt.Errorf("new db metric hook error: %w", err)
	}

	driverName, err := otsql.Register(
		DriverName,
		otsql.WithHooks(
			trace.New(
				trace.WithAllowRoot(true),
				trace.WithQuery(true),
				trace.WithQueryParams(true),
			),
			metricHook,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("register db metrics wrapper error, %w", err)
	}

	var db *sqlx.DB
	for i := 0; i < cfg.MaxConnectionAttempts+1; i++ {
		if i == cfg.MaxConnectionAttempts {
			return nil, fmt.Errorf("can't connect to db instance, %w", err)
		}

		db, err = sqlx.Connect(driverName, cfg.DSN)
		if err == nil {
			break
		}
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return db, nil
}
