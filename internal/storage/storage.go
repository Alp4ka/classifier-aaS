package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	// DB returns *sqlx.DB instance specified in NewStorage args.
	DB() *sqlx.DB

	// DBTx returns DBTx interface. It uses *sqlx.Tx if context contains it, otherwise it calls DB method.
	// It also sets isTx=true if transaction is used.
	DBTx(ctx context.Context) (dbTx DBTx, isTx bool)

	// GoquDBTx works the same way as Storage.DBTx method but wraps result into GoquDBTx interface.
	GoquDBTx(ctx context.Context) (goquDBTx GoquDBTx, isTx bool)

	// WithTransaction wraps scenario in function argument into a single transaction.
	//
	// Tries to fetch an existing transaction from context.
	// If the context did not contain transaction inside, it calls storage.db.Beginx() in order to create a new one and put
	// it into the context.
	WithTransaction(ctx context.Context, fn func(context.Context) error) error

	// Dialect returns database dialect. F.e. 'postgres'.
	Dialect() string
}
