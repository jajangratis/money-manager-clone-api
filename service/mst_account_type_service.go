package service

import (
	"context"
	"github.com/jajangratis/money-manager-clone-api/model/domain"
	"github.com/jajangratis/money-manager-clone-api/model/web"
)

type MstAccountTypeService interface {
	FindAll(ctx context.Context) []web.MstAccountTypeResponse
	Delete(ctx context.Context, id domain.MstAccountType)
	Create(ctx context.Context, request web.MstAccountCreateRequest) web.MstAccountTypeResponse
}
