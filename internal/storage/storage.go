package storage

import (
	"context"
	sqlpkg "github.com/Alp4ka/classifier-aaS/pkg/sql"
	"github.com/jmoiron/sqlx"
)

type SQLStorage interface {
	mustEmbedSQLStorageImpl()
	DBTX(ctx context.Context) sqlpkg.DBTX
}

type SQLStorageImpl struct {
	DB *sqlx.DB
}

func (s *SQLStorageImpl) mustEmbedSQLStorageImpl() {}

func (s *SQLStorageImpl) DBTX(ctx context.Context) sqlpkg.DBTX {
	tx, err := sqlpkg.TxFromCtx(ctx)
	if err != nil {
		return s.DB
	}

	return tx
}
