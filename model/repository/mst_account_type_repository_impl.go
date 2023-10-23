package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jajangratis/money-manager-clone-api/helper"
	"github.com/jajangratis/money-manager-clone-api/model/domain"
)

type MstAccountRepositoryImpl struct{}

func NewMstAccountRepositoryImpl() *MstAccountRepositoryImpl {
	return &MstAccountRepositoryImpl{}
}

func (repository *MstAccountRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MstAccountType {
	SQL := "SELECT * from mst_account_type"
	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	helper.PanicIfError(err)
	var mstAccountTypes []domain.MstAccountType
	for rows.Next() {
		accountType := domain.MstAccountType{}
		//parsedCreatedDate, err := time.Parse(time.RFC3339, &category.CreatedDate)
		//parsedUpdatedDate, err := time.Parse(time.RFC3339, &category.CreatedDate)
		err := rows.Scan(
			&accountType.Id,
			&accountType.AccountId,
			&accountType.AccountName,
			&accountType.CreatedDate,
			&accountType.UpdatedDate,
		)
		helper.PanicIfError(err)
		mstAccountTypes = append(mstAccountTypes, accountType)
	}
	return mstAccountTypes
}

func (repository *MstAccountRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, accountType domain.MstAccountType) {
	SQL := "DELETE from mst_account_type where id = $1"
	_, err := tx.ExecContext(ctx, SQL, accountType.Id)
	helper.PanicIfError(err)
}

func (repository *MstAccountRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, accountType domain.MstAccountType) domain.MstAccountType {
	SQL := "Insert into mst_account_type (account_id,account_name) values ($1, $2)"
	fmt.Println(SQL)
	_, err := tx.ExecContext(ctx, SQL, accountType.AccountId, accountType.AccountName)
	helper.PanicIfError(err)

	return accountType
}
