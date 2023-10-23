package repository

import (
	"context"
	"database/sql"
	"github.com/jajangratis/money-manager-clone-api/model/domain"
)

type MstAccountTypeRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MstAccountType
	Delete(ctx context.Context, tx *sql.Tx, accountType domain.MstAccountType)
	Save(ctx context.Context, tx *sql.Tx, accountType domain.MstAccountType) domain.MstAccountType
}
