package helper

import (
	"github.com/jajangratis/money-manager-clone-api/model/domain"
	"github.com/jajangratis/money-manager-clone-api/model/web"
)

func ToMstCategoryResponse(mstCategory domain.MstCategory) web.MstCategoryResponse {
	return web.MstCategoryResponse{
		Id:           mstCategory.Id,
		CategoryId:   mstCategory.CategoryId,
		CategoryName: mstCategory.CategoryName,
		CreatedDate:  mstCategory.CreatedDate,
		UpdatedDate:  mstCategory.UpdatedDate,
	}
}

func ToMstCategoryResponses(mstCategories []domain.MstCategory) []web.MstCategoryResponse {
	var mstCategoryResponses []web.MstCategoryResponse
	for _, mstCategory := range mstCategories {
		mstCategoryResponses = append(mstCategoryResponses, ToMstCategoryResponse(mstCategory))
	}
	return mstCategoryResponses
}

func ToMstAccountTypeResponse(mstCategory domain.MstAccountType) web.MstAccountTypeResponse {
	return web.MstAccountTypeResponse{
		Id:          mstCategory.Id,
		AccountId:   mstCategory.AccountId,
		AccountName: mstCategory.AccountName,
		CreatedDate: mstCategory.CreatedDate,
		UpdatedDate: mstCategory.UpdatedDate,
	}
}

func ToMstAccountTypeResponses(mstCategories []domain.MstAccountType) []web.MstAccountTypeResponse {
	var mstCategoryResponses []web.MstAccountTypeResponse
	for _, mstCategory := range mstCategories {
		mstCategoryResponses = append(mstCategoryResponses, ToMstAccountTypeResponse(mstCategory))
	}
	return mstCategoryResponses
}

func ToMstTransactionMethodResponse(mstTransactionMethod domain.MstTransactionMethod) web.MstTransactionMethodResponse {
	return web.MstTransactionMethodResponse{
		Id:          mstTransactionMethod.MethodId,
		MethodId:    mstTransactionMethod.MethodId,
		MethodName:  mstTransactionMethod.MethodName,
		CreatedDate: mstTransactionMethod.CreatedDate,
		UpdatedDate: mstTransactionMethod.UpdatedDate,
	}
}

func ToMstTransactionMethodResponses(mstTransactionMethods []domain.MstTransactionMethod) []web.MstTransactionMethodResponse {
	var mstTransactionMethodResponses []web.MstTransactionMethodResponse
	for _, mstCategory := range mstTransactionMethods {
		mstTransactionMethodResponses = append(mstTransactionMethodResponses, ToMstTransactionMethodResponse(mstCategory))
	}
	return mstTransactionMethodResponses
}
