package service

import (
	"context"
	"github.com/jajangratis/money-manager-clone-api/model/domain"
	"github.com/jajangratis/money-manager-clone-api/model/web"
)

type MstCategoryService interface {
	FindAll(ctx context.Context) []web.MstCategoryResponse
	Delete(ctx context.Context, mstCategory domain.MstCategory)
	Create(ctx context.Context, request web.MstCategoryCreateRequest) web.MstCategoryResponse
}
