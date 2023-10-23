package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jajangratis/money-manager-clone-api/helper"
	"github.com/jajangratis/money-manager-clone-api/model/domain"
)

type MstCategoryRepositoryImpl struct{}

func NewMstCategoryRepositoryImpl() *MstCategoryRepositoryImpl {
	return &MstCategoryRepositoryImpl{}
}

func (repository *MstCategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MstCategory {
	SQL := "SELECT * from mst_category"
	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	helper.PanicIfError(err)
	var mstCategories []domain.MstCategory
	for rows.Next() {
		category := domain.MstCategory{}
		//parsedCreatedDate, err := time.Parse(time.RFC3339, &category.CreatedDate)
		//parsedUpdatedDate, err := time.Parse(time.RFC3339, &category.CreatedDate)
		err := rows.Scan(
			&category.Id,
			&category.CategoryId,
			&category.CategoryName,
			&category.CreatedDate,
			&category.UpdatedDate,
		)
		helper.PanicIfError(err)
		mstCategories = append(mstCategories, category)
	}
	return mstCategories
}

func (repository *MstCategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, mstCategory domain.MstCategory) {
	SQL := "DELETE from mst_category where id = $1"
	_, err := tx.ExecContext(ctx, SQL, mstCategory.Id)
	helper.PanicIfError(err)
}

func (repository *MstCategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, mstCategory domain.MstCategory) domain.MstCategory {
	SQL := "Insert into mst_category (category_id,account_name) values ($1, $2)"
	fmt.Println(SQL)
	_, err := tx.ExecContext(ctx, SQL, mstCategory.CategoryId, mstCategory.CategoryName)
	helper.PanicIfError(err)
	return mstCategory
}
