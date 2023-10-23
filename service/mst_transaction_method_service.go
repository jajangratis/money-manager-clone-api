package service

import (
	"context"
	"github.com/jajangratis/money-manager-clone-api/model/domain"
	"github.com/jajangratis/money-manager-clone-api/model/web"
)

type MstTransactionMethodService interface {
	FindAll(ctx context.Context) []web.MstTransactionMethodResponse
	Delete(ctx context.Context, id domain.MstTransactionMethod)
	Create(ctx context.Context, request web.MstTransactionMethodRequest) web.MstTransactionMethodResponse
}
