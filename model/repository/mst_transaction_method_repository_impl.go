package repository

import (
	"context"
	"database/sql"
	"github.com/jajangratis/money-manager-clone-api/helper"
	"github.com/jajangratis/money-manager-clone-api/model/domain"
)

type MstTransactionMethodRepositoryImpl struct{}

func NewMstTransactionMethodRepositoryImpl() *MstTransactionMethodRepositoryImpl {
	return &MstTransactionMethodRepositoryImpl{}
}

func (repository *MstTransactionMethodRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MstTransactionMethod {
	SQL := "SELECT * from mst_transaction_method"
	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	helper.PanicIfError(err)
	var mstTransactionMethodTypes []domain.MstTransactionMethod
	for rows.Next() {
		mstTransactionMethodType := domain.MstTransactionMethod{}
		//parsedCreatedDate, err := time.Parse(time.RFC3339, &category.CreatedDate)
		//parsedUpdatedDate, err := time.Parse(time.RFC3339, &category.CreatedDate)
		err := rows.Scan(
			&mstTransactionMethodType.Id,
			&mstTransactionMethodType.MethodId,
			&mstTransactionMethodType.MethodName,
			&mstTransactionMethodType.CreatedDate,
			&mstTransactionMethodType.UpdatedDate,
		)
		helper.PanicIfError(err)
		mstTransactionMethodTypes = append(mstTransactionMethodTypes, mstTransactionMethodType)
	}
	return mstTransactionMethodTypes
}

func (repository *MstTransactionMethodRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, mstTransactionMethod domain.MstTransactionMethod) {
	SQL := "DELETE from mst_transaction_method where id = ?"
	_, err := tx.ExecContext(ctx, SQL, mstTransactionMethod.Id)
	helper.PanicIfError(err)
}

func (repository *MstTransactionMethodRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, mstTransactionMethod domain.MstTransactionMethod) domain.MstTransactionMethod {
	SQL := "Insert into mst_transaction_method(method_id, method_name) values ($1, $2)"
	_, err := tx.ExecContext(ctx, SQL, mstTransactionMethod.MethodId, mstTransactionMethod.MethodName)
	helper.PanicIfError(err)

	return mstTransactionMethod
}
