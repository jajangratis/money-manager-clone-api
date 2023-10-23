package repository

import (
	"context"
	"database/sql"
	"github.com/jajangratis/money-manager-clone-api/model/domain"
)

type MstTransactionMethodRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MstTransactionMethod
	Delete(ctx context.Context, tx *sql.Tx, mstTransactionMethod domain.MstTransactionMethod)
	Save(ctx context.Context, tx *sql.Tx, mstTransactionMethod domain.MstTransactionMethod) domain.MstTransactionMethod
}
