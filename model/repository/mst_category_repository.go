package repository

import (
	"context"
	"database/sql"
	"github.com/jajangratis/money-manager-clone-api/model/domain"
)

type MstCategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MstCategory
	Delete(ctx context.Context, tx *sql.Tx, mstCategory domain.MstCategory)
	Save(ctx context.Context, tx *sql.Tx, mstCategory domain.MstCategory) domain.MstCategory
}
